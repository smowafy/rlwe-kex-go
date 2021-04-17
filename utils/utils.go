package utils

import(
	"crypto/rand"
)

// interface mostly for mocks in tests
type RandomGenerator interface {
	Read([]byte) (int, error)
}

type CryptoRandomGen struct{}

func (rg *CryptoRandomGen) Read(b []byte) (int, error) {
	return rand.Read(b)
}

func NewRandomGenerator() RandomGenerator {
	return &CryptoRandomGen{}
}

func RandomBit() int {
	v := make([]byte, 1)

	_, err := rand.Read(v)

	if err != nil {
		panic(err)
	}

	return int(v[0]) & 1
}

func Log2(ln int) int {
	log2a := 0

	for {
		ln >>= 1

		if ln <= 0 {
			break
		}

		log2a++
	}

	return log2a
}

func RandomUInt32(rg RandomGenerator) uint32 {
	var ans uint32

	b := make([]byte, 4)
	_, err := rg.Read(b)

	if err != nil {
		panic(err)
	}

	for j := 0; j < 4; j++ {
		ans |= uint32(b[j]) << (j << 3)
	}

	return ans
}

func RandomInt64(rg RandomGenerator) int64 {
	var ans int64

	b := make([]byte, 8)
	_, err := rg.Read(b)

	if err != nil {
		panic(err)
	}

	for j := 0; j < 8; j++ {
		ans |= int64(b[j]) << (j << 3)
	}

	return ans
}

// 1 if a < b, otherwise 0
// a ^ ((a ^ b) | (a - b) ^ a)
// taken from the original repo:
// 	https://github.com/dstebila/rlwekex/blob/master/rlwe.c
func BitwiseLt(a, b int64) int64 {
	return (a ^ ((a ^ b) | (a - b) ^ a)) >> 63 & 1
}

// also taken from the repo
func BitwiseEq(a, b int64) int64 {
	return ((a - b) | (b - a)) >> 63
}

func BitwiseGtOrEqual(a, b int64) int64 {
	return BitwiseLt(b, a)
}

