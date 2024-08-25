package conn

import (
	"crypto/ed25519"
	"errors"
	"net"
)

const (
	CLIENT_HELLO  = "CLIENT_HELLO"
	SERVER_HELLO  = "SERVER_HELLO"
	X25519KeySize = 32
)

type request struct {
	Hello     string            `json:"hello"`
	PublicKey ed25519.PublicKey `json:"public_key"`
	Options   map[string]string `json:"options"`
}

func NewRequest(hello string, publicKey ed25519.PublicKey, options map[string]string) (*request, error) {
	if hello != CLIENT_HELLO {
		return nil, errors.New("failed to new request:\n\thello message expected to be CLIENT_HELLO")
	}
	if len(publicKey) != X25519KeySize {
		return nil, errors.New("failed to new request:\n\tpublicKey is not a valid X25519 public key")
	}
	return &request{
		Hello:     hello,
		PublicKey: publicKey,
		Options:   options,
	}, nil
}

func (r *request) SendForReply(conn net.Conn) (*reply, error) {

}


type reply struct {
	Hello     string            `json:"hello"`
	PublicKey ed25519.PublicKey `json:"public_key"`
	Options   map[string]string `json:"options"`
}

func NewReply(hello string, publicKey ed25519.PublicKey, options map[string]string) (*reply, error) {
	if hello != SERVER_HELLO {
		return nil, errors.New("failed to new request:\n\thello message expected to be SERVER_HELLO")
	}
	if len(publicKey) != X25519KeySize {
		return nil, errors.New("failed to new request:\n\tpublicKey is not a valid X25519 public key")
	}
	return &reply{
		Hello:     hello,
		PublicKey: publicKey,
		Options:   options,
	}, nil
}