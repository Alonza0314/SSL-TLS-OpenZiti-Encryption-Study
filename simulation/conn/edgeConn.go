package conn

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net"
	"simulation/config"
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
	rx       []byte
	tx       []byte
	txHeader []byte
	sender   pki.Encryptor
	receiver pki.Decryptor
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
		// keyPair, err := pki.NewKeyPair(viper.GetString("client.privateKey"), viper.GetString("client.publicKey"))
		// if err != nil {
		// 	return nil, errors.New("failed to new keyPair:\n\t" + err.Error())
		// }
		// ec.keyPair = keyPair

		if ec.keyPair, err = pki.NewKeyPair(); err != nil {
			return nil, errors.New("failed to new keyPair:\n\t" + err.Error())
		}
	} else {
		ec.keyPair = nil
	}

	ec.txHeader = make([]byte, config.STREAM_HEADER_SIZE)
	if _, err := rand.Read(ec.txHeader); err != nil {
		return nil, errors.New("failed to make txHeader:\n\t" + err.Error())
	}

	return &ec, nil
}

func (c *edgeConn) Connect() error {
	options := map[string]string{
		config.TX_HEADER: string(c.txHeader),
	}
	req, err := NewRequest(config.CLIENT_HELLO, c.keyPair.Public(), options)
	if err != nil {
		return errors.New("failed to connect:\n\t" + err.Error())
	}
	rep, err := req.SendForReply(c.conn)
	if err != nil {
		return errors.New("failed to get reply:\n\t" + err.Error())
	}
	
	if c.rx, c.tx, err = c.keyPair.ClientSessionKeys(rep.PublicKey); err != nil {
		return errors.New("failed to compute rx tx :\n\t" + err.Error())
	}

	if c.sender, err = pki.NewEncryptor(c.tx, c.txHeader); err != nil {
		return errors.New("failed to new encryptor\n\t%s\n" + err.Error())
	}

	if c.receiver, err = pki.NewDecryptor(c.rx, []byte(rep.Options[config.TX_HEADER])); err != nil {
		return errors.New("failed to new decryptor\n\t%s\n" + err.Error())
	}
	return nil
}

func (c *edgeConn) Communicate() error {
	// TODO
	// use graceful stop
	return nil
}

func (c *edgeConn) read(msg *[]byte) (int, error) {
	return 0, nil
}

func (c *edgeConn) write(msg []byte) (int, error) {
	return 0, nil
}

func (c *edgeConn) Close() error {
	return c.conn.Close()
}
