package conn

import (
	"errors"
	"fmt"
	"net"
	"simulation/pki"

	"github.com/spf13/viper"
)

type edgeConn struct {
	protocol string
	addr     string
	port     int
	conn     net.Conn
	crypto   bool
	keyPair  *pki.KeyPair
	rx       any
	tx       any
	sender   any
	receiver any
}

func NewEdgeConn(clinetCfg string) (*edgeConn, error) {
	var err error
	viper.SetConfigFile(clinetCfg)
	if err = viper.ReadInConfig(); err != nil {
		return nil, errors.New("failed to read config:\n\t" + err.Error())
	}

	ec := edgeConn{}
	ec.protocol = viper.GetString("client.protocol")
	ec.addr = viper.GetString("client.addr")
	ec.port = viper.GetInt("client.port")

	ec.conn, err = net.Dial(ec.protocol, fmt.Sprintf("%s:%v", ec.addr, ec.port))
	if err != nil {
		return nil, errors.New("failed to dail:\n\t" + err.Error())
	}

	ec.crypto = viper.IsSet("client.crypto")
	if ec.crypto {
		keyPair, err := pki.NewKeyPair(viper.GetString("client.privateKey"), viper.GetString("client.publicKey"))
		if err != nil {
			return nil, errors.New("failed to new keyPair:\n\t" + err.Error())
		}
		ec.keyPair = keyPair
	} else {
		ec.keyPair = nil
	}

	return &ec, nil
}

func (c *edgeConn) Connect() error {
	req, err := NewRequest(CLIENT_HELLO, c.keyPair.Public(), nil)
	if err != nil {
		return errors.New("failed to connect:\n\t" + err.Error())
	}
	rep, err := req.SendForReply(c.conn)
	if err != nil {
		return errors.New("failed to get reply:\n\t" + err.Error())
	}
	// TODO: compute with reply
	fmt.Println(rep)
	return nil
}

func (c *edgeConn) Write(msg []byte) (int, error) {
	return 0, nil
}
