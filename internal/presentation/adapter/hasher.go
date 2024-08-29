package adapter

import (
	"golang.org/x/crypto/argon2"
)

type Hasher interface {
	GeneareHash(password, salt string) string
}

type Argon2Hasher struct {
	saltLength uint32
	time       uint32
	memory     uint32
	threads    uint8
	keyLength  uint32
}

func NewHasher() *Argon2Hasher {
	return &Argon2Hasher{
		saltLength: 24,
		time:       2,
		memory:     64,
		threads:    4,
		keyLength:  32,
	}
}

func (h *Argon2Hasher) GenerateHash(password, salt []byte) string {
	hash := argon2.Key(password, salt, h.time, h.memory, h.threads, h.keyLength)
	return string(hash)
}
