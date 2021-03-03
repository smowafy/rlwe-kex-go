package main

import (
	"testing"
	"math"
	"math/rand"
)

func TestEnsureLongMod(t *testing.T) {
	var i, count int64

	for i = 0; i < 10000000; i++ {
		expectedRes := i % (int64(1)<<33 - 2)
		res := EnsureLongMod(i)
		if res != expectedRes {
			t.Errorf("long mod failed, expected %v but got %v\n", expectedRes, res)
		}

		count++
	}

	for i = 0; i < (1<<60); i+=rand.Int63n(1<<40)+(1<<30) {
		expectedRes := i % (int64(1)<<33 - 2)
		res := EnsureLongMod(i)
		if res != expectedRes {
			t.Errorf("long mod failed, expected %v but got %v\n", expectedRes, res)
		}

		count++
	}

	t.Logf("test cases count: %v\n", count)
}

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


func TestModularRound(t *testing.T) {
	longPolynomial := LongPolynomial{}

	longPolynomial.Coefficients[0] = int64(NumMod)
	longPolynomial.Coefficients[1] = int64(NumMod) + 1
	longPolynomial.Coefficients[2] = 2

	expectedRes := Polynomial{}

	res0 := math.Round(2.0 * float64(NumMod) / float64(LongMod))
	res1 := math.Round(2.0 * (float64(NumMod)+1) / float64(LongMod))
	res2 := math.Round(2.0 * 2 / float64(LongMod))

	expectedRes.Coefficients[0] = uint32(res0)
	expectedRes.Coefficients[1] = uint32(res1)
	expectedRes.Coefficients[2] = uint32(res2)

	res := longPolynomial.ModularRound()

	if *res != expectedRes {
		t.Errorf("polynomial modular round failed, expected %v but got %v\n", expectedRes, res)
	}
}

func TestCrossRound(t *testing.T) {
	longPolynomial := LongPolynomial{}

	longPolynomial.Coefficients[0] = CrossRoundQ4 - 5
	longPolynomial.Coefficients[1] = CrossRoundQ4
	longPolynomial.Coefficients[2] = CrossRoundQ4 + 5
	longPolynomial.Coefficients[3] = CrossRoundQ2 - 5
	longPolynomial.Coefficients[4] = CrossRoundQ2
	longPolynomial.Coefficients[5] = CrossRoundQ2 + 5
	longPolynomial.Coefficients[6] = CrossRoundQ2 + CrossRoundQ4 - 5
	longPolynomial.Coefficients[7] = CrossRoundQ2 + CrossRoundQ4
	longPolynomial.Coefficients[8] = CrossRoundQ2 + CrossRoundQ4 + 5

	expectedRes := Polynomial{}
	expectedRes.Coefficients[0] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[0]) / float64(LongMod))) % 2
	expectedRes.Coefficients[1] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[1]) / float64(LongMod))) % 2
	expectedRes.Coefficients[2] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[2]) / float64(LongMod))) % 2
	expectedRes.Coefficients[3] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[3]) / float64(LongMod))) % 2
	expectedRes.Coefficients[4] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[4]) / float64(LongMod))) % 2
	expectedRes.Coefficients[5] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[5]) / float64(LongMod))) % 2
	expectedRes.Coefficients[6] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[6]) / float64(LongMod))) % 2
	expectedRes.Coefficients[7] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[7]) / float64(LongMod))) % 2
	expectedRes.Coefficients[8] = uint32(math.Floor(4.0 * float64(longPolynomial.Coefficients[8]) / float64(LongMod))) % 2

	res := longPolynomial.CrossRound()

	if *res != expectedRes {
		t.Errorf("polynomial modular round failed, expected %v but got %v\n", expectedRes, res)
	}
}
