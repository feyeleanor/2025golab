package main

import (
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

				// there's more to reversing a string than just reversing the order of the
				// runes but for our purposes we'll assume it's plain ASCII
				// see https://github.com/shomali11/util/blob/master/xstrings/xstrings.go
				SendMessage(c, string(Reverse([]rune(s))))
			})
		})
	} else {
		log.Fatal(e)
	}
}
