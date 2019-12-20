using System.Net;
using System.Collections.Generic;
using Newtonsoft.Json;

namespace ExcaliburNetwork.Protocol
{

    public class Msg
    {
        public enum Code : int
        {
            Invalid = 0,
            GetUserInfo = 1,
            EnterStage,
            ExitStage,
            BattleStart,
            BattleEnd,
        }

        public enum Err : int
        {
            // 协议code服务端尚未实现
            ErrNotImplemented = 1,

            // 本次请求发送失败
            ErrSendFailure,

            // uid不存在
            ErrNotExist,

            // 造成db错误
            ErrDBError,

            // 协议json解析错误
            ErrDecodeJSON,

            // 从客户端的输入校验服务器现有的数据，不一致
            ErrOutOfSync,

        }


        public static Response<T> FromJson<T>(string json)
        {
            return JsonConvert.DeserializeObject<Response<T>>(json);
        }

        public static object FromJson(string json,System.Type type)
        {
            return JsonConvert.DeserializeObject(json,type);
        }

        public virtual string ToJson()
        {
            return JsonConvert.SerializeObject(this);
        }

        public override string ToString()
        {
            return this.ToJson();
        }

        public static object MakeErrorResponse(Code code, long uid, int err, string errMsg = null)
        {
            switch (code)
            {
                case Msg.Code.GetUserInfo:
                    return new Response<GetUserInfo> { uid = uid, err = err, errMsg = errMsg  };
                case Msg.Code.EnterStage:
                    return new Response<EnterStage> { uid = uid, err = err, errMsg = errMsg };
                case Msg.Code.ExitStage:
                    return new Response<ExitStage> { uid = uid, err = err, errMsg = errMsg };
                case Msg.Code.BattleStart:
                    return new Response<BattleStart> { uid = uid, err = err, errMsg = errMsg };
                case Msg.Code.BattleEnd:
                    return new Response<BattleEnd> { uid = uid, err = err, errMsg = errMsg };
                default:
                    return null;
            }
        }
    }

    public class Response<T> : Msg
    {
        public Msg.Code code;
        public long uid;
        public int seq;
        public int err;
        public string errMsg = "";
        public T data;
    }


    public class GetUserInfoRequest
    {
    }

    public class GetUserInfo
    {
        public User user;
    }

    public class EnterStageRequest
    {
        public string characterId;
        public StageInfo stageInfo;
    }

    public class EnterStage
    {
        public MapInfo mapInfo;
        public Attribute battleAttribute;
    }

    public class ExitStageRequest
    {
        public string characterId;
        public StageInfo stageInfo;
        public bool stageClear;
    }

    public class ExitStage
    {
        public Bag battleBag;
    }

    public class BattleStartRequest
    {
        public string characterId;
        public int npcId;
        public int npcType;
    }

    public class BattleStart
    {
        public Attribute battleAttribute;
    }

    public class BattleEndRequest
    {
        public string characterId;
        public int npcId;
        public int npcType;
        public string battleResult;
        public Attribute battleAttribute;
    }

    public class BattleEnd
    {
    }
}
