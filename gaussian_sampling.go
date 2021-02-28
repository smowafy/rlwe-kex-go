package main

import (
	"crypto/rand"
)

type Vec struct {
	Chunks [3]uint64
}

// interface mostly for mocks in tests
type RandomGenerator interface {
	Read([]byte) (int, error)
}

type CryptoRandomGen struct { }

func (rg *CryptoRandomGen) Read(b []byte) (int, error) {
	return rand.Read(b)
}

// table of the precomputed values of the probability distribution described in
// the paper (J. W. Bos, C. Costello, M. Naehrig and D. Stebila, "Post-Quantum
// Key Exchange for the TLS Protocol from the Ring Learning with Errors
// Problem," 2015 IEEE Symposium on Security and Privacy, San Jose, CA, USA,
// 2015, pp. 553-570, doi: 10.1109/SP.2015.40.), which is almost identical to a
// discrete gaussian distribution over Z with mean 0 and standard deviation
// (8 / sqrt(2 * PI))

var table = [...]Vec {
	Vec { Chunks: [3]uint64 { 0x1FFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF } },
	Vec { Chunks: [3]uint64 { 0x5CEF2C248806C827, 0x161ABD186DA13542, 0xE0C81DA0D6A8BD22 } },
	Vec { Chunks: [3]uint64 { 0x9186506BCC065F20, 0x4344C125B3533F22, 0x8D026C4E14BC7408 } },
	Vec { Chunks: [3]uint64 { 0xBAAB5F82BCDB43B3, 0x5D62CE65E6217813, 0x10AC7CEC7D7E2A3B } },
	Vec { Chunks: [3]uint64 { 0xD7D9769FAD23BCB1, 0x1411F551608E4D22, 0x709C92996E94D801 } },
	Vec { Chunks: [3]uint64 { 0xEA9BE2F4D6DDB5ED, 0x7E1526D618902F20, 0x6287D827008404B7 } },
	Vec { Chunks: [3]uint64 { 0xF58A99474919B8C9, 0xE7D2A13787E94674, 0x34CBDC118C15F40E } },
	Vec { Chunks: [3]uint64 { 0xFB5117812753B7B8, 0xE8A773D9A1EA0AAB, 0xD521F7EBBBE8C3A2 } },
	Vec { Chunks: [3]uint64 { 0xFE151BD0928596D3, 0x148CB49FF716491B, 0xC3D9E58131089A6A } },
	Vec { Chunks: [3]uint64 { 0xFF487508BA9F7208, 0x07E44D009ADB0049, 0x2E060C4A842A27F6 } },
	Vec { Chunks: [3]uint64 { 0xFFC16686270CFC82, 0x1A5409BF5D4B039E, 0xFCEDEFCFAA887582 } },
	Vec { Chunks: [3]uint64 { 0xFFEC8AC3C159431B, 0xFDC99BFE0F991958, 0x4FE22E5DF9FAAC20 } },
	Vec { Chunks: [3]uint64 { 0xFFFA7DF4B6E92C28, 0xA6FCD4C13F4AFCE0, 0xA36605F81B14FEDF } },
	Vec { Chunks: [3]uint64 { 0xFFFE94BB4554B5AC, 0x4B869C6286ED0BB5, 0x9D1FDCFF97BBC957 } },
	Vec { Chunks: [3]uint64 { 0xFFFFAADE1B1CAA95, 0xEC72329E974D63C7, 0x6B3EEBA74AAD104B } },
	Vec { Chunks: [3]uint64 { 0xFFFFEDDC1C6436DC, 0x337F6316C1FF0A59, 0x48C8DA4009C10760 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFC7C9DC2569A, 0xD95E7B2CD6933C97, 0x84480A71312F35E7 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFF61BC337FED, 0x8E0B132AE72F729F, 0x23C01DAC1513FA0F } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFE6B3CF05F7, 0x05B9D725AAEA5CAD, 0x90C89D6570165907 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFC53EA610E, 0x99E8F72C370F27A6, 0x692E2A94C500EC7D } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFF841943DE, 0xC6E2F0D7CAFA9AB8, 0x28C2998CEAE37CC8 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFF12D07EC, 0x4745913CB4F9E4DD, 0xC515CF4CB0130256 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFE63E348, 0xEE62D42142AC6544, 0x39F0ECEA047D6E3A } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFD762C7, 0x064A0C6CC136E943, 0xDF11BB25B50462D6 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFC5E37, 0xC672F3A74DB0F175, 0xCDBA0DD69FD2EA0F } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFB48F, 0x6ABEF8B144723D83, 0xFDB966A75F3604D9 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFA72, 0x697598CEADD71A15, 0x3C4FECBB600740D1 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFA1, 0x12F5A30DD99D7051, 0x1574CC916D60E673 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFA, 0x4016ED3E05883572, 0xDD3DCD1B9CB7321D } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xAF22D9AFAD5A73CF, 0xB4A4E8CF3DF79A7A } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFBF88681905332BA, 0x91056A8196F74466 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFD16385AF29A51F, 0x965B9ED9BD366C04 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFE16FF8EA2B60C, 0xF05F75D38F2D28A3 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFEDD3C9DDC7E8, 0x77E35C8980421EE8 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFF63392B6E8F, 0x92783617956F140A } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFB3592B3D1, 0xA536DC994639AD78 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFDE04A5BB, 0x8F3A871874DD9FD5 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFF257152, 0x310DE3650170B717 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFB057B, 0x1F21A853A422F8CC } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFE5AD, 0x3CA9D5C6DB4EE2BA } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFF81, 0xCFD9CE958E59869C } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFD, 0xDB8E1F91D955C452 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xF78EE3A8E99E08C3 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFE1D7858BABDA25 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFF9E52E32CAB4A } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFEE13217574F } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFD04888041 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFF8CD8A56 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFF04111 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFE0C5 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFC7 } },
	Vec { Chunks: [3]uint64 { 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF, 0xFFFFFFFFFFFFFFFF } },
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

	return int(v[0])&1
}

func (v *Vec) addWithCarry(v1 Vec, carry uint64) *Vec {
	var sum uint64 = 0

	for i := 2; i >= 0; i-- {
		sum = v.Chunks[i] + v1.Chunks[i] + carry

		// TODO: remove the if condition, side channel potential
		if sum < v.Chunks[i] {
			carry = 1
		} else {
			carry = 0
		}

		v.Chunks[i] = sum
	}

	return v
}

func NewVec(chunks [3]uint64) *Vec {
	return &Vec { Chunks: chunks }
}

func NewZeroVec() *Vec {
	return &Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0x0000000000000000,
			0x0000000000000000,
		},
	}
}

func NewOneVec() *Vec {
	return &Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0x0000000000000000,
			0x0000000000000001,
		},
	}
}

func NewRandomVec(rg RandomGenerator) *Vec {
	b := make([]byte, 24)
	_, err := rg.Read(b)

	if err != nil {
		panic(err)
	}

	var chunks [3]uint64

	for i := 0; i < 3; i++ {
		for j := 0; j < 8; j++ {
			chunks[i] |= uint64(b[(i<<3)+j]) << (j<<3)
		}
	}

	return &Vec {
		Chunks: chunks,
	}
}

func (v *Vec) Copy() *Vec {
	return &Vec { Chunks: v.Chunks }
}

// public
func (v *Vec) Xor(v1 Vec) *Vec {
	for i := 2; i >= 0; i-- {
		v.Chunks[i] = v.Chunks[i] ^ v1.Chunks[i]
	}

	return v
}

func (v *Vec) Or(v1 Vec) *Vec {
	for i := 2; i >= 0; i-- {
		v.Chunks[i] = v.Chunks[i] | v1.Chunks[i]
	}

	return v
}


func (v *Vec) And(v1 Vec) *Vec {
	for i := 2; i >= 0; i-- {
		v.Chunks[i] = v.Chunks[i] & v1.Chunks[i]
	}

	return v
}

func (v* Vec) Flip() *Vec {
	for i := 2; i >= 0; i-- {
		v.Chunks[i] = ^v.Chunks[i]
	}

	return v
}

func (v* Vec) Neg() *Vec {
	return v.Flip().addWithCarry(*NewZeroVec(), uint64(1))
}

// x ^ ((x ^ y) | (x - y) ^ x)
func (v Vec) BitwiseLt(v1 Vec) int {
	a := v.Copy().Xor(v1)
	b := v.Copy().Sub(v1)
	b.Xor(v)
	b.Or(*a)
	b.Xor(v)

	return int((b.Chunks[0]>>63) & 1)
}


func (v *Vec) Add(v1 Vec) *Vec {
	return v.addWithCarry(v1, uint64(0))
}

func (v *Vec) Sub(v1 Vec) *Vec {
	return v.addWithCarry(*(v1.Copy().Flip()), uint64(1))
}


func GaussianOverZ() int {
	vec := NewRandomVec(NewRandomGenerator())

	// 1 if the bit is 1
	// -1 if the bit is 0
	randSign := (RandomBit()<<1) - 1

	return constantTimeSampling(*vec) * randSign
}

func constantTimeSampling(randomVec Vec) int {
	mask, res := 0, 0

	for i := 0; i < 52; i++ {
		compRes := randomVec.BitwiseLt(table[i])

		// either 0 or -1
		newMask := -compRes
		// also either 0 or -1
		tmpMask := mask ^ newMask

		res |= (i & tmpMask)

		mask = newMask
	}

	return res
}
