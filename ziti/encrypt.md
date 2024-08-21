# Encrypt

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

```go
    // secretstream/stream.go
    func NewEncryptor(key []byte) (Encryptor, []byte, error) {
        if len(key) != StreamKeyBytes {
            return nil, nil, invalidKey
        }

        header := make([]byte, StreamHeaderBytes)
        _, err := rand.Read(header)
        if err != nil {
            return nil, nil, err
        }

        stream := &encryptor{}

        k, err := chacha20.HChaCha20(key[:], header[:16])
        if err != nil {
            return nil, nil, err
        }
        copy(stream.k[:], k)
        stream.reset()

        for i := range stream.pad {
            stream.pad[i] = 0
        }

        for i, b := range header[crypto_core_hchacha20_INPUTBYTES:] {
            stream.nonce[i+crypto_secretstream_xchacha20poly1305_COUNTERBYTES] = b
        }

        return stream, header, nil
    }

    func (s *encryptor) Push(plain []byte, tag byte) ([]byte, error) {
        var err error
        var poly *poly1305.MAC
        var block [64]byte
        var slen [8]byte

        mlen := len(plain)
        out := make([]byte, mlen+StreamABytes)

        chacha, err := chacha20.NewUnauthenticatedCipher(s.k[:], s.nonce[:])
        if err != nil {
            return nil, err
        }
        
        chacha.XORKeyStream(block[:], block[:])

        var poly_init [32]byte
        copy(poly_init[:], block[:])
        poly = poly1305.New(&poly_init)

        memzero(block[:])
        block[0] = tag

        chacha.XORKeyStream(block[:], block[:])
        _, _ = poly.Write(block[:])
        out[0] = block[0]

        c := out[1:]
        chacha.XORKeyStream(c, plain)
        _, _ = poly.Write(c[:mlen])
        padlen := (0x10 - len(block) + mlen) & 0xf
        _, _ = poly.Write(pad0[:padlen])

        binary.LittleEndian.PutUint64(slen[:], uint64(0))
        _, _ = poly.Write(slen[:])

        binary.LittleEndian.PutUint64(slen[:], uint64(len(block)+mlen))
        _, _ = poly.Write(slen[:])

        mac := c[mlen:]
        copy(mac, poly.Sum(nil))

        xor_buf(s.nonce[crypto_secretstream_xchacha20poly1305_COUNTERBYTES:], mac)
        buf_inc(s.nonce[:crypto_secretstream_xchacha20poly1305_COUNTERBYTES])

        return out, nil
    }
```
