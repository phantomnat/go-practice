package main

import (
	crand "crypto/rand"
	"math/big"
	"math/rand"
	"testing"
)

func BenchmarkRandMath(b *testing.B) {
	for n := 0; n < b.N; n++ {
		rand.Int63()
	}
}

func BenchmarkRandCrypto(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := crand.Int(crand.Reader, big.NewInt(27))
		if err != nil {
			panic(err)
		}
	}
}
