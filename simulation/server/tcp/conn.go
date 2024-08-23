package tcp

import (
	"fmt"
	"log"
	"net"
)

type ListenerInterface func(string, int, HandlerInterface)

type HandlerInterface func(conn net.Conn)

func TCPListener(host string, port int, handler HandlerInterface) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Fatalf("Failed to listen on %s:%v: %s\n", host, port, err.Error())
	}
}

func TCPHandler(conn net.Conn) {
}
