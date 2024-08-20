# Handshake

## Source Code(Handshake)

```go
    conn, err := context.Dial(serviceName)
```

```go
    // sdk-golang/ziti/ziti.go
    func (context *ContextImpl) Dial(serviceName string) (edge.Conn, error) {
        // ~~~
        return context.DialWithOptions(serviceName, defaultOptions)
    }
```

```go
    // sdk-golang/ziti/ziti.go
    func (context *ContextImpl) DialWithOptions(serviceName string, options *DialOptions) (edge.Conn, error) {
        // ~~~
        conn, err := context.dialSession(svc, session, edgeDialOptions)
        //~~~
    }
```

```go
    // sdk-golang/ziti/ziti.go
    func (context *ContextImpl) dialSession(service *rest_model.ServiceDetail, session *rest_model.SessionDetail, options *edge.DialOptions) (edge.Conn, error) {
        // ~~~
        return edgeConnFactory.Connect(service, session, options)
    }
```

```go
    // sdk-golang/ziti/edge/network/conn.go
    func (conn *edgeConn) Connect(session *rest_model.SessionDetail, options *edge.DialOptions) (edge.Conn, error) {
        // ~~~
        conn.establishClientCrypto(conn.keyPair, hostPubKey, edge.CryptoMethod(method))
        // ~~~
    }
```

```go
    // sdk-golang/ziti/edge/network/conn.go
    func (conn *edgeConn) establishClientCrypto(keypair *kx.KeyPair, peerKey []byte, method edge.CryptoMethod) error {
        var err error
        var rx, tx []byte

        if method != edge.CryptoMethodLibsodium {
            return unsupportedCrypto
        }

        if rx, tx, err = keypair.ClientSessionKeys(peerKey); err != nil {
            return errors.Wrap(err, "failed key exchange")
        }

        var txHeader []byte
        if conn.sender, txHeader, err = secretstream.NewEncryptor(tx); err != nil {
            return errors.Wrap(err, "failed to establish crypto stream")
        }

        conn.rxKey = rx

        if _, err = conn.MsgChannel.Write(txHeader); err != nil {
            return errors.Wrap(err, "failed to write crypto header")
        }

        pfxlog.Logger().
            WithField("connId", conn.Id()).
            WithField("marker", conn.marker).
            Debug("crypto established")
        return nil
    }
```
