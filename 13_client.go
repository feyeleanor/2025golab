package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("needs name of domain socket to create")
	}

	DialServer("unix", os.Args[1], func(c net.Conn) {
		if len(os.Args) > 2 {
			s, _ := json.Marshal(os.Args[2:])
			SendMessage(c, s)
			if m, e := ReceiveMessage(c); e == nil {
				r := []any{}
				if e = json.Unmarshal([]byte(m), &r); e == nil {
					for i, v := range r {
						log.Printf("sent [%v], received [%v]", os.Args[i+2], v)
					}
				} else {
					log.Println(e)
				}
			} else {
				log.Println(e)
			}
		}
	})
}
