package conn

import "time"

const (
	CLIENT_HELLO  = "CLIENT_HELLO"
	SERVER_HELLO  = "SERVER_HELLO"
	X25519KEYSIZE = 32
	BUFFERSIZE    = 1024
	CONN_TIMEOUT  = 5 * time.Second

	NEW      = "[NEW]"
	RECEIVED = "[RECEIVED]"
	SENT     = "[SENT]"
)
