package msg

// Msg 消息结构
type Request struct {
	Code Code   `json:"code"`
	Uid  int64  `json:"uid"`
	Seq  int    `json:"seq"`
	Data string `json:"data"`
}

type Response struct {
	Code   Code        `json:"code"`
	Uid    int64       `json:"uid"`
	Seq    int         `json:"seq"`
	Err    int         `json:"err"`
	ErrMsg string      `json:"errMsg"`
	Data   interface{} `json:"data"`
}

func FromRequest(req *Request) *Response {
	return &Response{
		Code: req.Code,
		Uid:  req.Uid,
		Seq:  req.Seq,
	}
}

func OkFromRequest(req *Request, data interface{}) *Response {
	return &Response{
		Code: req.Code,
		Uid:  req.Uid,
		Seq:  req.Seq,
		Data: data,
	}
}

func ErrFromRequest(req *Request, err int, errmsg string) *Response {
	return &Response{
		Code:   req.Code,
		Uid:    req.Uid,
		Seq:    req.Seq,
		Err:    err,
		ErrMsg: errmsg,
	}
}

func OK(req *Request, data interface{}) *Response {
	return &Response{
		Code: req.Code,
		Uid:  req.Uid,
		Seq:  req.Seq,
		Data: data,
	}
}

func Err(req *Request, err int, errmsg string) *Response {
	return &Response{
		Code:   req.Code,
		Uid:    req.Uid,
		Seq:    req.Seq,
		Err:    err,
		ErrMsg: errmsg,
	}
}

type ErrCode = int

type IMsgBase struct {
	errcode ErrCode
}

func (b IMsgBase) ErrCode() ErrCode {
	return b.errcode
}
