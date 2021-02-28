package main

import(
	"testing"
)

func equal(a, b [3]uint64) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestVecAdd(t *testing.T) {
	v1 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0xFFFFFFFFFFFFFFFF,
			0xFFFFFFFFFFFFFFFE,
		},
	}

	v2 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0x0000000000000000,
			0x0000000000000002,
		},
	}

	v3 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000001,
			0x0000000000000000,
			0x0000000000000000,
		},
	}

	v1.Add(v2)

	for i := 0; i < 3; i++ {
		if(v1.Chunks[i] != v3.Chunks[i]) {
			t.Errorf("Sum not equal, got %v and %v\n", v1, v3)
		}
	}
}

func TestVecSub(t *testing.T) {
	v1 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000001,
			0x0000000000000000,
			0x0000000000000000,
		},
	}

	v2 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0x0000000000000000,
			0x0000000000000002,
		},
	}

	v3 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0xFFFFFFFFFFFFFFFF,
			0xFFFFFFFFFFFFFFFE,
		},
	}

	v1.Sub(v2)

	for i := 0; i < 3; i++ {
		if(v1.Chunks[i] != v3.Chunks[i]) {
			t.Errorf("Diff not equal, got %v and %v\n", v1, v3)
		}
	}
}

func TestVecCopy(t *testing.T) {
	v1 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000001,
			0x0000000000000000,
			0x0000000000000000,
		},
	}

	v2 := v1.Copy()

	v2.Chunks[0] = 0x0000000000000000

	if v1.Chunks == v2.Chunks {
		t.Errorf("pointers are similar, %v %v\n", v1.Chunks, v2.Chunks)
	}
}

func TestVecBitwiseLt(t *testing.T) {
	v1 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000000,
			0x000000000000FFFF,
			0x0000000000000000,
		},
	}

	v2 := Vec {
		Chunks: [3]uint64 {
			0x0000000000000001,
			0x0000000000000000,
			0x0000000000000000,
		},
	}


	if v1.BitwiseLt(v2) != 1 {
		t.Errorf("answer  = %v\n", v1.BitwiseLt(v2))
	}
}

// https://medium.com/@ankur_anand/how-to-mock-in-your-go-golang-tests-b9eee7d7c266
var fakeRead func(b []byte) (int, error)

type fakeRandomGenerator struct { }

func (rg fakeRandomGenerator) Read(b []byte) (int, error) {
	return fakeRead(b)
}

func TestNewRandomVec(t *testing.T) {
	var expectedChunks = [3]uint64 {
		0xFE151BD0928596D3, 0x148CB49FF716491B, 0xC3D9E58131089A6A,
	}

	fakeRead = func (b []byte) (int, error) {
		for i := 0; i < 3; i++ {
			for j := 0; j < 8; j++ {
				b[(i<<3) + j] = byte((expectedChunks[i]>>(j<<3)) & 0xFF)
			}
		}

		return 24, nil
	}

	randomVec := NewRandomVec(fakeRandomGenerator{})

	if !equal(randomVec.Chunks, expectedChunks) {
		t.Errorf(
			"Random vector not equal to the expected one, values: %v, %v\n",
			randomVec.Chunks,
			expectedChunks,
		)
	}
}

func TestConstsantTimeSampling(t *testing.T) {
	for i := 0; i < 1; i++ {
		res := constantTimeSampling(table[i])
		if res != i+1 {
			t.Errorf("failure in sampling, %v didn't produce the expected index %v but produced %v\n", table[i], i+1, res)
		}

		res = constantTimeSampling(table[i])
		if res != i+1 {
			t.Errorf("failure in sampling, %v didn't produce the expected index %v but produced %v\n", table[i], i+1, res)
		}
	}
}
