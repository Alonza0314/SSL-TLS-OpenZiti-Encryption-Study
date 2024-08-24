package conn

import (
	"crypto/ed25519"
	"net"
)

var CLIENT_HELLO string = "clientHello"
var SERVER_HELLO string = "serverHello"

type Request struct {
	hello     string
	publicKey ed25519.PublicKey
	options   map[string]string
}
type Reply struct {
	hello     string
	publicKey ed25519.PublicKey
	options   map[string]string
}

func NewRequest() (*Request, error) {

}

func NewReply() (*Reply, error) {
	
}

func (r *Request) SendForReply(conn net.Conn) []byte {

}
