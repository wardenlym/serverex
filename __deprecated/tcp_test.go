package main

import (
	"fmt"
	"net"
)

// a simple test for pure tcp
func test() {
	addr, _ := net.ResolveTCPAddr("tcp", ":8082")
	listen, _ := net.ListenTCP("tcp", addr)
	go func() {
		for {
			conn, e := listen.AcceptTCP()
			fmt.Println(conn.RemoteAddr(), e)

			go func() {
				defer fmt.Println(conn.RemoteAddr().String() + "out")
				for {
					var e error
					buf := make([]byte, 1024)
					_, e = conn.Read(buf)
					if e != nil {
						return
					}
					s := "ok,you are" + conn.RemoteAddr().String() + string(buf)
					_, e = conn.Write([]byte(s))
					if e != nil {
						return
					}
				}
			}()
		}
	}()
}
