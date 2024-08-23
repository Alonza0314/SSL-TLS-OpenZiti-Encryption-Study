package main

import "server/tcp"

func main() {
	var _ tcp.ListenerInterface = tcp.TCPListener
	var _ tcp.HandlerInterface = tcp.TCPHandler
	tcp.TCPListener("localhost", 8080, tcp.TCPHandler)
}