# TLS in Ziti

The model seems like TLS 1.3.

## [Handshake](handshake.md)

### Connect

1. In connect function, we first check if it needs "crypto" and extract its(local) public key.
2. And, we send the first "hello" message including public key and some configs to peer endpoint.
3. Then we get the crypto method from the reply message, including TLS version and encrypt alogorithm.
4. Also, we get the peer side public key from reply message.
5. Then, it's time to estblish client crypto, i.e., call the "establishClientCrypto" function.

### establishClientCrypto

1. Check if it is based on "libsodium"; otherwise, break it.
2. Use local key and peer key and the "ClientSessionKeys" function to make the session key: rx(decrypt received data), tx(encrypt sent data).
3. Create an "encryptor" as "conn.sender" which is used to encrypt and send data.
4. Then, we need to send "txHeader" to peer side for secretstream initialization.
5. Finally, we will use this sender to send data and use rx to decrypt message.

## [Encrypt](encrypt.md)

---

## Reference

[Source code: sdk-golang](https://github.com/openziti/sdk-golang)

[ChatGPT](https://openai.com/chatgpt/)
