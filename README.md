# ssl-tls-Research

## Index

[Abstract](#abstract)

[Basic Concept](#basic-concept)

[Encryption Method](#encryption-method)

[SSL](#ssl)

[TLS](#tls)

[Reference](#reference)

## Abstract

Introduce the historical development of SSL and TLS and their significance in network security.

## Basic Concept

### Introduction

+ SSL(Secure Sockets Layer): An early encryption protocol developed by Netscape, which has now been replaced by TLS.
+ TLS (Transport Layer Security): The successor to SSL, providing a more secure encryption protocol.

### Key Function

+ Encryption: Ensures that data is not intercepted or altered during transmission.
+ Authentication: Ensures the trustworthiness of the communicating parties.
+ Data Integrity: Ensures that data is not modified during transmission.

## Encryption Method

### Symmetric and Asymmetric Encryption

+ Symmetric Encryption: Uses the same key for both encryption and decryption (e.g., DES \ AES).
    ![symmetricEncryption](static/img/symmetricEncyrption.avif)
+ Asymmetric Encryption: Uses a pair of public and private keys for encryption and decryption (e.g., RSA).
    ![asymmetricEncryption](static/img/asymmetricEncryption.avif)

## SSL

### SSL 1.0

+ Description: Never publicly released due to serious security issues.

### SSL 2.0

+ Improvements:
  + Introduced basic encryption mechanisms.
Supported symmetric encryption and digital certificates.
+ Vulnerabilities:
  + Several known security vulnerabilities, such as insecure key exchange mechanisms.
  + Did not support Message Authentication Codes (MACs) to verify data integrity.

### SSL 3.0

+ Improvements:
  + Addressed multiple security issues from SSL 2.0, improving encryption algorithms and protocol design.
  + Introduced a more secure handshake process.
  + Supported Message Authentication Codes (MACs), enhancing data integrity protection.
+ Vulnerabilities:
  + Despite improvements in security, SSL 3.0 still had some issues, such as the [POODLE attack](https://www.acunetix.com/blog/web-security-zone/what-is-poodle-attack/) (Padding Oracle On Downgraded Legacy Encryption).

## TLS

### TLS 1.0

### TLS 1.1

### TLS 1.2

### TLS 1.3

---

## Reference

[Transport Layer Security](https://en.wikipedia.org/wiki/Transport_Layer_Security)

[[Secure101] 從歷史到分級， 一篇文章看懂什麼是 SSL](https://simular.co/blog/post/21-%E4%B8%80%E7%AF%87%E6%96%87%E7%AB%A0%E7%9C%8B%E6%87%82%E4%BB%80%E9%BA%BC%E6%98%AFssl)

[SSL/TLS and PKI History](https://www.feistyduck.com/ssl-tls-and-pki-history/)