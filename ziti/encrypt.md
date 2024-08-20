# Encryption

```go
    conn.Write([]byte(""))
```

```go
    // sdk-golang/ziti/edge/network/conn.go
    func (conn *edgeConn) Write(data []byte) (int, error) {
        if conn.sentFIN.Load() {
            return 0, errors.New("calling Write() after CloseWrite()")
        }

        if conn.sender != nil {
            cipherData, err := conn.sender.Push(data, secretstream.TagMessage)
            if err != nil {
                return 0, err
            }

            _, err = conn.MsgChannel.Write(cipherData)
            return len(data), err
        } else {
            return conn.MsgChannel.Write(data)
        }
    }
```
