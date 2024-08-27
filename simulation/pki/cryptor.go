package pki

import (
	"errors"
	"fmt"
	"simulation/config"

	"golang.org/x/crypto/chacha20"
	"golang.org/x/crypto/chacha20poly1305"
)

var pad0 [16]byte

type streamState struct {
	k     [config.STREAM_KEY_SIZE]byte
	nonce [chacha20poly1305.NonceSize]byte
	pad   [8]byte
}

func (s *streamState) reset() {
	for i := range s.nonce {
		s.nonce[i] = 0
	}
	s.nonce[0] = 1
}

type encryptor struct {
	streamState
}

type Encryptor interface {
	Push()
}

func NewEncryptor(key, header []byte) (Encryptor, error) {
	if len(key) != config.STREAM_KEY_SIZE {
		return nil, errors.New("failed to new encryptor:\n\tkey length is expected to be " + fmt.Sprint(config.STREAM_KEY_SIZE) + " but " + fmt.Sprint(len(key)))
	}

	stream := &encryptor{}

	k, err := chacha20.HChaCha20(key, header[:16])
	if err != nil {
		return nil, errors.New("failed to new encryptor:\n\t" + err.Error())
	}
	copy(stream.k[:], k)
	stream.reset()

	for i := range stream.pad {
		stream.pad[i] = 0
	}

	for i, b := range header[config.CRYPTO_CORE_HCHACHA20_INPUTSIZE:] {
		stream.nonce[i+config.CRYPTO_SECRETSTREAM_XCHACHA20POLY1305_COUNTERBYTES] = b
	}

	return stream, nil
}

func (e *encryptor) Push() {

}

type decryptor struct {
	streamState
}

type Decryptor interface {
	Pull()
}

func NewDecryptor(key, header []byte) (Decryptor, error) {
	stream := &decryptor{}
	k, err := chacha20.HChaCha20(key, header[:config.CRYPTO_CORE_HCHACHA20_INPUTSIZE])
	if err != nil {
		return nil, errors.New("failed to new decryptor:\n\t" + err.Error())
	}
	copy(stream.k[:], k)

	stream.reset()

	copy(stream.nonce[config.CRYPTO_SECRETSTREAM_XCHACHA20POLY1305_COUNTERBYTES:], header[config.CRYPTO_CORE_HCHACHA20_INPUTSIZE:])

	copy(stream.pad[:], pad0[:])

	return stream, nil
}

func (d *decryptor) Pull() {
	
}
