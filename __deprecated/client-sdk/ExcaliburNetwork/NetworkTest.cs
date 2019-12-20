using System.Collections;
using System.Collections.Generic;
using UnityEngine;

using System.Linq;
using Newtonsoft.Json;
using ExcaliburNetwork;
using ExcaliburNetwork.Protocol;

class ExampleCallback : ClientCallback
{
    public override void OnEvent(object o)
    {

    }
    public override void OnResponse(Msg.Code code, object o ,object ctx)
    {
        Debug.Log("callback OnResponse: "+code+" "+o);
        switch (code)
        {
            case Msg.Code.GetUserInfo:
                var response = Msg.FromJson<GetUserInfo>(o.ToString());
                break;
            case Msg.Code.EnterStage:
                var r2 = Msg.FromJson<EnterStage>(o.ToString());
                Debug.Log(r2.ToJson());
                break;
            default:
                Debug.Log("default callback OnResponse: " + code + " " + o);
                break;
        }
    }
    public override void OnError(object o)
    {
        Debug.Log("callback OnError: "+o);
    }
}

public class NetworkTest : MonoBehaviour {

    Client client;

	// Use this for initialization
	void Start () {
        Debug.Log(System.Environment.Version); // 2.0
        Debug.Log("start");

        client = new Client(ClientOption.DefaultBetaOption, Debug.unityLogger);
        client.SetClientCallback(new ExampleCallback());

        client.Start();

        // client.Login()
        // 稍后登录做完了会取到uid,暂时随便设置
        client.UserId = 666;


        // 协议传输数据类
        Response<GetUserInfo> info = new Response<GetUserInfo>();
        string json = info.ToJson();
        Debug.Log(json);
        var info2 = Msg.FromJson<GetUserInfo>(json);
        Debug.Log(info2.ToString());

        foreach (var i in Enumerable.Range(0, 5))
        {
            // do something
            client.SendGameRequest(Msg.Code.GetUserInfo,new GetUserInfoRequest());
        }

        // 验证性能测试: client端，大约 平均收发 16次/每秒
        // curl - i - XPOST 'http://api.lambda-studio.cn/excalibur/v1/game?uid=123&code=1' - w "\n%{time_namelookup}\n%{time_connect}\n%{time_starttransfer}\n%{time_total}"

        // 战斗相关

        var enterReq = new EnterStageRequest();
        enterReq.characterId = "1";
        enterReq.stageInfo = new StageInfo { chapterId = 1, stageId = 1 };
        client.SendGameRequest(Msg.Code.EnterStage, enterReq);
        client.SendGameRequest(Msg.Code.BattleStart, new BattleStartRequest { characterId = "1", npcId = 1, npcType = 1 });

        // battleResult 等待客户端定义：打过？没打过？逃跑？
        client.SendGameRequest(Msg.Code.BattleEnd, new BattleEndRequest { characterId = "1", npcId = 1, npcType = 1, battleResult = "没打过", battleAttribute = new ExcaliburNetwork.Protocol.Attribute { hp = 998 } });
        client.SendGameRequest(Msg.Code.ExitStage, new ExitStageRequest());

        // 可以看到数据变了 battleAttribute.hp=998 lastStageInfo=1,1
        client.SendGameRequest(Msg.Code.GetUserInfo);

    }
	
	// Update is called once per frame
	void Update () {
		
	}

    private void OnDestroy()
    {
        if (client != null)
        {
            client.Stop();
            Debug.Log("exit");
        }
    }
}
