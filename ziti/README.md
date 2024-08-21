# TLS in Ziti

## [Handshake](handshake.md)

### Connect

    1. In connect function, we first check if it needs "crypto" and extract its(client's) public key.
    2. Then we get the crypto method, including TLS version and encrypt alogorithm.
    3. Also, we get the server public key from reply message.
    4. Then, it's time to estblish client crypto, i.e., call the "establishClientCrypto" function.

### establishClientCrypto

    1. 

## [Encrypt](encrypt.md)
