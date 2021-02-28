package main

import (
	"testing"
	"math/rand"
)

func TestModAdd(t *testing.T) {
	var a, b, res uint32
	var expectedRes uint64

	const cases int = 1000

	for c := 0; c < cases; c++ {
		a, b = rand.Uint32(), rand.Uint32()

		expectedRes = (uint64(a) + uint64(b)) % uint64(NumMod)

		res = ModAdd(a, b)

		if EnsureMod(uint32(expectedRes)) != res {
			t.Errorf("mod addition failed, expected %v but got %v\n", expectedRes, res)
		}
	}
}

func TestModMultiply(t *testing.T) {
	var a, b, res uint32
	var expectedRes uint64

	const cases int = 1000

	for c := 0; c < cases; c++ {
		a, b = rand.Uint32(), rand.Uint32()

		expectedRes = (uint64(a) * uint64(b)) % (uint64(NumMod))

		res = ModMultiply(a, b)

		if EnsureMod(uint32(expectedRes)) != res {
			t.Errorf("mod addition failed, expected %v but got %v\n", expectedRes, res)
		}
	}
}

func TestModNeg(t *testing.T) {
	var a, b uint32

	a = 17
	b = ModNeg(a)

	if ModAdd(a, b) != 0 {
		t.Errorf("mod negation modulo %v failed, %v + %v != 0 but %v\n", NumMod, a, b, ModAdd(a, b))
	}

	if 0 != ModNeg(0) {
		t.Errorf("mod negation modulo %v failed, ModNeg(0) != 0 but %v\n", NumMod, ModNeg(0))
	}
}

func TestPolyAdd(t *testing.T) {
}

func TestPolyMultiply(t *testing.T) {
	p1, p2, expectedRes := Polynomial{}, Polynomial{}, Polynomial{}

	p1.Coefficients[650] = 1
	p1.Coefficients[100] = 9
	p1.Coefficients[2] = ModNeg(2)
	p1.Coefficients[1] = 1
	p1.Coefficients[0] = 4

	p2.Coefficients[650] = 7
	p2.Coefficients[0] = 1

	expectedRes.Coefficients[750] = 63
	expectedRes.Coefficients[652] = ModNeg(14)
	expectedRes.Coefficients[651] = 7
	expectedRes.Coefficients[650] = 29
	expectedRes.Coefficients[276] = ModNeg(7)
	expectedRes.Coefficients[100] = 9
	expectedRes.Coefficients[2] = ModNeg(2)
	expectedRes.Coefficients[1] = 1
	expectedRes.Coefficients[0] = 4

	if p1.Multiply(&p2).Coefficients != expectedRes.Coefficients {
		t.Errorf("polynomial multiplication failed, expected %v but got %v\n", expectedRes, p1)
	}
}
