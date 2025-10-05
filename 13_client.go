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
				ProcessJSON(m, func(i int, v string) {
					log.Printf("sent [%v], received [%v]", os.Args[i+2], v)
				})
			} else {
				log.Println(e)
			}
		}
	})
}

func ProcessJSON[T any](b []byte, f func(int, T)) {
	m := []any{}
	if e := json.Unmarshal(b, &m); e == nil {
		log.Println(m)
		for i, v := range m {
			if v, ok := v.(T); ok {
				f(i, v)
			}
		}
	} else {
		log.Println(e)
	}
}
