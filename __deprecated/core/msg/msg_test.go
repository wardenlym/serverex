package msg

import (
	"fmt"
	"testing"
)

func Test_Code(t *testing.T) {
	fmt.Println(C_Any)
	fmt.Println(C_Multicast)

	fmt.Println(C_Invalid)

	fmt.Println(C_AgentConnect)
	fmt.Println(C_AgentDisconnect)
	fmt.Println(C_AgentInactive)
	fmt.Println(C_AgentDeactive)
}
