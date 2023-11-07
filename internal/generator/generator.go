package generator

import (
	"github.com/ericlagergren/saferand"
	"time"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_"

type Generator struct{}

func NewGenerator() *Generator {
	return &Generator{}
}

func (g *Generator) GenerateHash() string {
	saferand.Seed(uint64(time.Now().UTC().UnixNano()))
	hash := make([]byte, 10)
	alphLen := len(alphabet)
	for i := range hash {
		hash[i] = alphabet[saferand.Intn(alphLen)]
	}
	return string(hash)
}
