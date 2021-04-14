package polynomial

import(
	"github.com/smowafy/rlwe-kex-go/gaussian"
)

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

func RandomUInt32(rg gaussian.RandomGenerator) uint32 {
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

func RandomInt64(rg gaussian.RandomGenerator) int64 {
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

// produces 0 if given 0 (NumMod + 1)
// produces 2^32 - 1 (NumMod) if given 1
func BitmaskFromBit(b uint32) uint32 {
	return NumMod + ((b & 1) ^ 1)
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

