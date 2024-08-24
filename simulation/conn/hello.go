package conn

import "crypto/ed25519"

var CLIENT_HELLO string = "clientHello"
var SERVER_HELLO string = "serverHello"

type Request struct {
	hello string
	publicKey	ed25519.PublicKey
	options map[string]string
}

