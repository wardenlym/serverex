package codegen

import (
	"io/ioutil"
	"regexp"
	"strings"
	"testing"
)

func Test_codegen_protocol_src(t *testing.T) {

	b, e := ioutil.ReadFile("protocol.def.go")
	if e != nil {
		t.Error(e)
	}

	interp1 := new_interp()
	for i, v := range strings.Split(string(b), "\n") {
		interp1.emit(i+1, v)
	}

	// fmt.Println(interp1.ret)

	interp1.complie()

	ioutil.WriteFile("../msg/msg.go", []byte(interp1.src_msg_golang), 0664)
	ioutil.WriteFile("./sdk/Excalibur/Protocol/Msg.cs", []byte(interp1.src_msg_csharp), 0644)
	ioutil.WriteFile("E:\\Excalibur\\Clent\\Assets\\Network\\sdk\\Excalibur\\Protocol\\Msg.cs", []byte(interp1.src_msg_csharp), 0644)

	ioutil.WriteFile("../odm/odm.go", []byte(interp1.src_odm_golang), 0644)
	ioutil.WriteFile("./sdk/Excalibur/Protocol/Odm.cs", []byte(interp1.src_odm_csharp), 0644)
	ioutil.WriteFile("E:\\Excalibur\\Clent\\Assets\\Network\\sdk\\Excalibur\\Protocol\\Odm.cs", []byte(interp1.src_odm_csharp), 0644)

	ioutil.WriteFile("../msg/request.go", []byte(interp1.src_request_golang), 0644)
	ioutil.WriteFile("./sdk/Excalibur/Protocol/Request.cs", []byte(interp1.src_request_csharp), 0644)
	ioutil.WriteFile("E:\\Excalibur\\Clent\\Assets\\Network\\sdk\\Excalibur\\Protocol\\Request.cs", []byte(interp1.src_request_csharp), 0644)

	ioutil.WriteFile("./sdk/Excalibur/Protocol/Rpc.cs", []byte(interp1.src_rpc_csharp), 0644)
	ioutil.WriteFile("E:\\Excalibur\\Clent\\Assets\\Network\\sdk\\Excalibur\\Protocol\\Rpc.cs", []byte(interp1.src_rpc_csharp), 0644)

}

func Test_reg(t *testing.T) {
	var enum_field_reg_with_comment = regexp.MustCompile(`^\s+(\w+)\s+//\s+(.+)$`)

	ss := enum_field_reg_with_comment.FindStringSubmatch("	   ErrForTest   //    iam comment")
	if len(ss) > 0 {

		println(ss[1], ss[2])
	}
}
