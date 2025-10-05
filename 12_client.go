package main

import (
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("needs name of domain socket to create")
	}

	DialServer("unix", os.Args[1], func(c net.Conn) {
		for _, n := range os.Args[2:] {
			SendMessage(c, n)
			if m, e := ReceiveMessage(c); e == nil {
				log.Printf("sent [%v], received [%v]", n, string(m))
			} else {
				log.Println(e)
			}
		}
	})
}
