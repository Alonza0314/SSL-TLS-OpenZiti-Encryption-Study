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
}

func NewEdgeConn(clinetCfg string) (*edgeConn, error) {
	var err error
	viper.SetConfigFile(clinetCfg)
	if err = viper.ReadInConfig(); err != nil {
		return nil, errors.New("failed to read config: " + err.Error())
	}

	ec := edgeConn{}
	ec.protocol = viper.GetString("client.protocol")
	ec.addr = viper.GetString("client.addr")
	ec.port = viper.GetInt("client.port")

	ec.conn, err = net.Dial(ec.protocol, fmt.Sprintf("%s:%v", ec.addr, ec.port))
	if err != nil {
		return nil, errors.New("failed to dail: " + err.Error())
	}

	ec.crypto = viper.IsSet("client.crypto")
	if ec.crypto {
		keyPair, err := pki.NewKeyPair(viper.GetString("client.privateKey"), viper.GetString("client.publicKey"))
		if err != nil {
			return nil, errors.New("failed to new keyPair: " + err.Error())
		}
		ec.keyPair = keyPair
	} else {
		ec.keyPair = nil
	}
	
	return &ec, nil
}

func (c *edgeConn) Connect() error {
	return nil
}

func (c *edgeConn) Write(msg []byte) (int, error) {
	return 0, nil
}
