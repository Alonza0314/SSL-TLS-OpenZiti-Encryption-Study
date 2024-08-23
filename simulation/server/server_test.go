package server

import (
	nt "server/network"
	"testing"
)

func TestServer(t *testing.T) {
	var _ nt.ListenerInterface = nt.TCPListener
	var _ nt.HandlerInterface = nt.TCPHandler
	nt.TCPListener("localhost", 8080, nt.TCPHandler)
}
