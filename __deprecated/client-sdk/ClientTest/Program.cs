using System;
using System.Threading;
using System.Collections;

namespace ClientTest
{
    using System.Linq;
    using ExcaliburNetwork.Protocol;
    using Newtonsoft.Json;
    using UnityEngine;

    class Program
    {
        class ExampleCallback : ExcaliburNetwork.ClientCallback
        {
            public override void OnEvent(object o)
            {

            }
            public override void OnResponse(Msg.Code code, object o, object ctx)
            {
                Console.WriteLine($"callback OnResponse: {code} {o}");
            }
            public override void OnError(object o)
            {
                Console.WriteLine($"callback OnError: {o}");
            }
        }

        static void Main(string[] args)
        {
            //int a, b;
            //ThreadPool.GetAvailableThreads(out a, out b);
            //Console.WriteLine("原有辅助线程:" + a + "原有I/O线程:" + b);

            var client = new ExcaliburNetwork.Client(ExcaliburNetwork.ClientOption.DefaultBetaOption, null);
            client.SetClientCallback(new ExampleCallback());
            client.Start();

            // 协议传输数据类
            Response<GetUserInfo> info = new Response<GetUserInfo>();
            string json = info.ToJson();
            Console.WriteLine(json);
            var info2 = Msg.FromJson<GetUserInfo>(json);
            Console.WriteLine(info2.ToString());

            // 发送Game消息,尝试10秒内收发200条
            //foreach (var i in Enumerable.Range(0, 10))
            //{
            //    // do something
            //    client.SendGameRequest(Msg.Code.GetUserInfo, i);
            //}

            // 验证性能测试: client端，大约 平均收发 16次/每秒
            // curl - i - XPOST 'http://api.lambda-studio.cn/excalibur/v1/game?uid=123&code=1' - w "\n%{time_namelookup}\n%{time_connect}\n%{time_starttransfer}\n%{time_total}"


            client.UserId = 123;
            client.SendGameRequest(Msg.Code.GetUserInfo);

            var enterReq = new EnterStageRequest();
            enterReq.characterId = "1";
            enterReq.stageInfo = new StageInfo { chapterId = 1, stageId = 1 };
            client.SendGameRequest(Msg.Code.EnterStage, enterReq);
            client.SendGameRequest(Msg.Code.BattleStart, new BattleStartRequest { characterId="1",npcId=1,npcType=1 });
            client.SendGameRequest(Msg.Code.BattleEnd, new BattleEndRequest { characterId = "1", npcId = 1, npcType = 1, battleResult = "没打过", battleAttribute = new Attribute { hp=998} });
            client.SendGameRequest(Msg.Code.ExitStage, new ExitStageRequest());


            client.SendGameRequest(Msg.Code.GetUserInfo);

            Thread.Sleep(10_000);
            client.Stop();
            Console.ReadKey();
        }
    }
}
