package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"

	"gitlab.com/megatech/serverex/core/pkt"
	"gitlab.com/megatech/serverex/lib/log"
	"gitlab.com/megatech/serverex/lib/util"
)

func main() {
	log.EnableDebug(false)
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:11002")
	if err != nil {
		log.Error("ResolveTCPAddr failed:", err)
		os.Exit(1)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	// Read from stdin, then send packet(type:Echo) to server
	outgoing := make(chan *pkt.Pkt)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for {
			for scanner.Scan() {
				pkt := &pkt.Pkt{
					Code: pkt.C_Echo,
					Data: []byte(scanner.Text()),
				}
				outgoing <- pkt
			}
			if scanner.Err() != nil {
				log.Fatal(scanner.Err())
			}
		}
	}()
	go func() {
		encoder := pkt.NewLengthPrefixPacker()
		for {
			packet := <-outgoing
			log.Debug(packet.Code, packet.Data)

			buf := bytes.NewBuffer([]byte{})
			encoder.Encode(packet, buf)

			n, err := conn.Write(buf.Bytes())
			if err != nil {
				if err == io.EOF {
					log.Info("remote peer closed", n, err)
					os.Exit(0)
				} else {
					log.Fatal(n, err)
				}
			}
		}
	}()

	// Recv packet from server, then display to stdin
	incoming := make(chan *pkt.Pkt)
	go func() {
		decoder := pkt.NewLengthPrefixPacker()
		for {
			pkt, err := decoder.Decode(conn)
			if err != nil {
				if err == io.EOF {
					log.Info("remote peer closed", err)
					os.Exit(0)
				} else {
					log.Fatal(err)
				}
			}
			incoming <- pkt
		}
	}()
	go func() {
		for {
			pkt, ok := <-incoming
			if pkt == nil {
				log.Error(pkt, ok)
				return
			}
			fmt.Printf("> [%v] %s\n", pkt.Code, string(pkt.Data))
		}
	}()

	<-util.WaitInterruptSignal()
	conn.Close()
	fmt.Printf("client exit.\n\n")
}
