# Excalibur网络层客户端sdk说明

(随sdk代码一起更新，不另行通知,update之后如果发现sdk的代码改动很多，记得查看一下这个文档)

### · 目前有联调过的只有GetUserInfo和EnterStage两个接口！！

### · 用新的uid调用接口前先确保调用过一次GetUserInfo，这样服务端才能生成这个用户的初始信息！！

- 错误定位方法:

调试过程中当出现跟你的预期不符时？

1. 如果接口有返回Response的情况下

    正确处理Response中的Err字段
    
    当Err!=0的情况下 查看 ExcaliburNetwork.Protocol下 Msg.Err

    ```
            public enum Err : int
        {
            // 协议code服务端尚未实现
            ErrNotImplemented = 1,

            // uid不存在
            ErrNotExist,

            // 造成db错误
            ErrDBError,

            // 协议json解析错误
            ErrDecodeJSON,

            // 从客户端的输入校验服务器现有的数据，不一致
            ErrOutOfSync,
        }
    ```

2. 接口有返回，但是错误仍然不明了

    打开 http://log.lambda-studio.cn 这是beta服务端的weblog

    第一次打开会默认输出上几次调用的日志

    如果没有你要的，就再调用几次接口，这个web页面会实时输出日志，不要频繁刷新，除非长时间没反应

    找到返回接口中的seq字段，这个流水号在每次请求时自动生成
    
    ```
    public class Res : Msg
    {
        public Msg.Code code;
        public long uid;
        public int seq;            // 调用序号
        public int err;
        public string errMsg = "";
    }
    ```

    对应到日志中 "-> #" "<- #" 后面的数字就是收发内容:

    比如下面这个可以浏览器里Ctrl-f 搜 #10,可以快速定位到问题,
    
    之后可以吧这个日志或序号告诉我,方便交流查bug

    ```
    
    2018-03-26 10:43:29.653784 DEBUG [global] [-> #10 C_ExitStage /excalibur/v1/game?uid=666&code=3 {"battleBag":null,"code":0,"uid":0,"seq":0,"err":0,"errMsg":""}] [gameplay.go gameplay.(*game).Handle(54)]

    2018-03-26 10:43:29.653943 DEBUG [global] [<- #10 C_ExitStage {"code":3,"uid":666,"seq":10,"err":0,"errMsg":""}] [gameplay.go gameplay.(*game).Handle(63)]

    ```

3. 接口超时，或不明原因没有返回

    服务有自动重启，程序内部有异常回复，所以一般说服务器不会不通的

    可以在浏览器里访问
    
    http://api.lambda-studio.cn/test

    返回

    ```
    {"status":"running"}
    ```

    或访问

    http://api.lambda-studio.cn/gateway 


    返回

    ```
    {"login":"http://api.lambda-studio.cn/excalibur/v1/account","game":"http://api.lambda-studio.cn/excalibur/v1/game"}
    ```

    都能说明服务器在正常运行中


4. fxxk，确实不通？！

    我在临时处理必须停机的问题，一般如果没有通知几分钟后重试就可以，长时间的会在群里通知

5. 无法理解的长时间不通

    只有一种可能，腾讯云故障。
    
    (其实还有一种可能，服务器没钱了。解决方法：通知gaoyang，给我打钱)