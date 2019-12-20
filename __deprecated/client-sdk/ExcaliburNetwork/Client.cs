namespace ExcaliburNetwork
{
    using System;
    using System.IO;
    using System.Net;
    using System.Threading;
    using System.Collections;
    using System.Text;
    using UnityEngine;
    using ExcaliburNetwork.Protocol;
    using System.Collections.Generic;
    using Newtonsoft.Json;

    public interface IClientCallback
    {
        void OnEvent(object o);
        void OnResponse(Msg.Code code, object o, object ctx);
        void OnError(object o);
    }
    public class ClientCallback : IClientCallback
    {
        public virtual void OnEvent(object o)
        {

        }
        public virtual void OnResponse(Msg.Code code, object o, object ctx)
        {
        }
        public virtual void OnError(object o)
        {
        }
    }

    public sealed class ClientOption
    {
        public string GatewayURL;
        
        // 连本机的开发用服务器,一般客户端不会用到
        public static ClientOption DefaultDevOption = new ClientOption
        {
            GatewayURL = "http://localhost:8080/gateway",
        };

        // 内部上线beta版的服务器地址,默认就用这个
        public static ClientOption DefaultBetaOption = new ClientOption
        {
            GatewayURL = "http://api.lambda-studio.cn:80/gateway",
        };
    }

    public sealed class Client
    {
        public const string Version = "0.0.1";
        public long UserId;
        string gameServerURL = "";
        string loginServerURL = "";
        IClientCallback callback = new ClientCallback();
        ILogger logger = null;
        public Client(ClientOption option = null, ILogger logger = null)
        {
            if (option != null)
            {
               this.option = option;
            }
            else
            {
                this.option = ClientOption.DefaultBetaOption;
            }
            
            if (logger != null)
            {
                this.logger = logger;
            }
        }

        void log(LogType type,object o)
        {
            if (this.logger != null)
            {
                logger.Log(type,o);
            }
            else
            {
                Console.WriteLine(type.ToString() + " " + o.ToString());
            }
        }

        public void SetClientCallback(IClientCallback callback)
        {
            this.callback = callback;
        }

        public bool Start()
        {
            // 重要：HttpWebRequest接口的一些调参
            // https://en.code-bude.net/2013/01/21/3-things-you-should-know-to-speed-up-httpwebrequest/
            // https://stackoverflow.com/questions/2519655/httpwebrequest-is-extremely-slow

            WebRequest.DefaultWebProxy = null;
            System.Net.ServicePointManager.Expect100Continue = false;
            ServicePointManager.DefaultConnectionLimit = 100;

            return init() && start();
        }

        public void Stop()
        {
            uninit();
        }

        public void SendGameRequest(Msg.Code code, object o = null, object ctx = null)
        {
            SendGameRequest(code, this.UserId, o, ctx);
        }

        public void SendGameRequest(Msg.Code code, long uid, object o = null, object ctx = null)
        {

            queue.Enqueue(()=>{
                Stream dataStreamSend = null;
                WebResponse response = null;
                StreamReader reader = null;
                Stream dataStreamRecv = null;
                try
                {
                    var url = gameServerURL +"?uid="+uid+"&code="+(int)code;

                    HttpWebRequest request = (HttpWebRequest)HttpWebRequest.Create(url);
                    request.Method = "POST";
                    request.Timeout = 3 * 1000;
                    request.ReadWriteTimeout = 3 * 1000;

                    byte[] byteArray = new byte[] { };
                    if (o != null)
                    {
                        byteArray = Encoding.UTF8.GetBytes(JsonConvert.SerializeObject(o));
                    }
                    request.ContentLength = byteArray.Length;
                    dataStreamSend = request.GetRequestStream();
                    dataStreamSend.Write(byteArray, 0, byteArray.Length);
                    dataStreamSend.Close();
                    response = request.GetResponse();
                    dataStreamRecv = response.GetResponseStream();
                    reader = new StreamReader(dataStreamRecv);
                    string responseFromServer = reader.ReadToEnd();
                    callback.OnResponse(code, responseFromServer,ctx);
                }
                catch (WebException ex)
                {
                    using (HttpWebResponse res = (HttpWebResponse)ex.Response)
                    {
                        switch (res.StatusCode)
                        {
                            case HttpStatusCode.BadRequest:
                            case HttpStatusCode.InternalServerError:
                                var errRet = Msg.MakeErrorResponse(code,uid, (int)Msg.Err.ErrSendFailure, res.ToString());

                                if (errRet == null)
                                {
                                    callback.OnError(ex);
                                }
                                else
                                {
                                    callback.OnResponse(code, errRet, ctx);
                                }
                                
                                break;
                            default:
                                callback.OnError(ex);
                                break;
                        }
                    }
                }
                catch (Exception ex)
                {
                    // shit happen ... 待处理
                    callback.OnError(ex);
                }
                finally
                {
                    if (dataStreamSend != null)
                    {
                        dataStreamSend.Close();
                    }
                    if (response != null)
                    {
                        response.Close();
                    }
                    if (reader != null)
                    {
                        reader.Close();
                    }
                    if (dataStreamRecv != null)
                    {
                        dataStreamRecv.Close();
                    }
                }
            });
        }

        void SendHttpRequestAsync(WebRequest request)
        {
            queue.Enqueue(execute(request));
        }

        [Obsolete("这个方法仅供测试不要在unity里使用")]
        public string SendHttpRequestBlocking(WebRequest request)
        {
            string param = "hl=zh-CN&newwindow=1";
            byte[] bs = Encoding.ASCII.GetBytes(param);
            request.ContentLength = bs.Length;

            string res;
            using (WebResponse wr = request.GetResponse())
            {
                using (var stream = wr.GetResponseStream())
                {
                    using (var reader = new StreamReader(stream, Encoding.UTF8))
                    {
                        res = reader.ReadToEnd();
                    }
                }
            }
            return res;
        }

        Thread thread;
        ClientOption option;
        BlockingQueue<Action> queue;
        bool running = false;

        bool init()
        {
            thread = new Thread(mainLoop);
            queue = new BlockingQueue<Action>();
            return true;
        }
        bool start()
        {
            running = true;
            thread.Start();
            beginGetGateway();
            return true;
        }
        void uninit()
        {
            running = false;
            queue.Enqueue(() => { });
            thread.Join();
        }


        void mainLoop()
        {
            while(running)
            {
                try
                {
                    var action = queue.Dequeue();
                    action();
                }
                catch (Exception ex)
                {
                    callback.OnError(ex);
                }
                
            }
        }

        Action execute(WebRequest request)
        {
            return () => { };
        }

        void beginGetGateway()
        {
            queue.Enqueue(() =>
            {
                WebResponse response = null;
                StreamReader reader = null;
                Stream dataStreamRecv = null;
                try
                {
                    var url = option.GatewayURL;
                    HttpWebRequest request = (HttpWebRequest)HttpWebRequest.Create(url);
                    request.Method = "GET";
                    response = request.GetResponse();
                    dataStreamRecv = response.GetResponseStream();
                    reader = new StreamReader(dataStreamRecv);
                    string responseFromServer = reader.ReadToEnd();
                    Console.WriteLine(responseFromServer);
                    var map = JsonConvert.DeserializeObject<Dictionary<string, string>>(responseFromServer);
                    this.gameServerURL = map["game"];
                }
                catch (WebException ex)
                {
                    using (HttpWebResponse res = (HttpWebResponse)ex.Response)
                    {
                        switch (res.StatusCode)
                        {
                            case HttpStatusCode.BadRequest:
                            case HttpStatusCode.InternalServerError:
                                callback.OnError(res.StatusCode);
                                break;
                            default:
                                Console.WriteLine(ex.Status + " " + ex + " ");
                                break;
                        }
                    }
                }
                catch (Exception ex)
                {
                    // shit happen ...
                    Console.WriteLine(ex);
                }
                finally
                {
                    if (response != null)
                    {
                        response.Close();
                    }
                    if (reader != null)
                    {
                        reader.Close();
                    }
                    if (dataStreamRecv != null)
                    {
                        dataStreamRecv.Close();
                    }
                }
            });
        }
    }
}
