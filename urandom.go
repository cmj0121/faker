package faker

import (
	"crypto/rand"
	"math"
	"math/big"
)

var (
	maxInt64 = big.NewInt(math.MaxInt64)
)

type CryptoRandom struct {
}

// crypto-random not support seed
func (crandom CryptoRandom) Seed(seed int64) {
	// NOP
	return
}

func (crandom CryptoRandom) Int63() (out int64) {
	if bigint, err := rand.Int(rand.Reader, maxInt64); err != nil {
		// cannot get crypto/rand
		panic(err)
	} else {
		// save as the int64
		out = bigint.Int64()
	}
	return
}

func (crandom CryptoRandom) Float64() (out float64) {
	if numerator, err := rand.Int(rand.Reader, maxInt64); err != nil {
		// cannot get crypto/rand
		panic(err)
	} else if denominator, err := rand.Int(rand.Reader, maxInt64); err != nil {
		// cannot get crypto/rand
		panic(err)
	} else {
		// save as the float64
		out, _ = big.NewRat(numerator.Int64(), denominator.Int64()).Float64()
	}
	return
}
