package config

import "time"

const (
	CLIENT_HELLO = "CLIENT_HELLO"
	SERVER_HELLO = "SERVER_HELLO"

	BUFFERSIZE = 1024

	X25519KEYSIZE  = 32
	SESSIONKEYSIZE = 32
	CONN_TIMEOUT   = 5 * time.Second

	NEW      = "[NEW]"
	RECEIVED = "[RECEIVED]"
	SENT     = "[SENT]"
)
