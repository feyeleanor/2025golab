package main

import (
	"encoding/json"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("needs name of domain socket to create")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove(os.Args[1])
		os.Exit(1)
	}()

	if s, e := net.Listen("unix", os.Args[1]); e == nil {
		OnConnection(s, func(c net.Conn) {
			MessageLoop(c, func(s string) {
				log.Println("Received:", s)
				m := TransformJSON(s, func(v string) string {
					return string(Reverse([]rune(v)))
				})
				SendMessage(c, m)
			})
		})
	} else {
		log.Fatal(e)
	}
}

func TransformJSON[T any](s string, f func(T) T) string {
	m := []any{}
	if e := json.Unmarshal([]byte(s), &m); e == nil {
		log.Println(m)
		for i, v := range m {
			if v, ok := v.(T); ok {
				m[i] = f(v)
			}
		}
	} else {
		log.Println(e)
	}

	r, _ := json.Marshal(m)
	return string(r)
}
