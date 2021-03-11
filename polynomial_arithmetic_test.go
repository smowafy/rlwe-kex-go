package main

import (
	"testing"
	"math"
	"math/rand"
)

func TestEnsureMod(t *testing.T) {
	var i, count uint32

	for i = 0; i < 10000000; i++ {
		expectedRes := uint32(int64(i) % (int64(1)<<33 - 2))
		res := EnsureMod(i)
		if res != expectedRes {
			t.Errorf("mod failed, expected %v but got %v\n", expectedRes, res)
		}

		count++
	}

	if EnsureMod(NumMod) != 0 {
		t.Errorf("mod failed, expected %v but got %v\n", 1, EnsureMod(NumMod))
	}

	for i = 0; i < (1<<31); i+=(rand.Uint32() % (1<<30)) {
		expectedRes := uint32(int64(i) % (int64(1)<<33 - 2))
		res := EnsureMod(i)
		if res != expectedRes {
			t.Errorf("mod failed, expected %v but got %v\n", expectedRes, res)
		}

		count++
	}

	t.Logf("test cases count: %v\n", count)
}

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

func TestAdjustPositiveNegative(t *testing.T) {
	cases := 1000

	var c int64

	for i := 0; i < cases; i++ {
		c = rand.Int63n(LongMod>>2)

		if AdjustPositiveNegative(LongModNeg(c)) != -c {
			t.Errorf("adjusting failed, expected %v but got %v\n", -c, AdjustPositiveNegative(LongModNeg(c)))
			return
		}
	}
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

func TestLongModAdd(t *testing.T) {
	var a, b, res int64
	var expectedRes int64

	const cases int = 1000

	for c := 0; c < cases; c++ {
		a, b = rand.Int63n((1<<62)), rand.Int63n((1<<62))

		expectedRes = (a + b) % LongMod

		res = LongModAdd(a, b)

		if EnsureLongMod(expectedRes) != res {
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

func TestLongModNeg(t *testing.T) {
	for a := (int64(1)<<32); a <= (int64(1)<<34); a+=(rand.Int63n(int64(1)<<20)) {
		expected := (-a + 8 * LongMod) % LongMod
		actual := LongModNeg(a)

		if expected != actual {
			t.Errorf("mod negation modulo %v failed for %v, expected %v but got %v\n", LongMod, a, expected, actual)

			return
		}
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

//	longPolynomial.Coefficients[0] = int64(NumMod)
	longPolynomial.Coefficients[0] = int64(NumMod>>2) - 1
	longPolynomial.Coefficients[1] = int64(NumMod) - 3
	longPolynomial.Coefficients[2] = 2

	expectedRes := Polynomial{}

	res0 := math.Round(2.0 * (float64(AdjustPositiveNegative(int64(NumMod>>2) - 1))) / float64(LongMod))
	res1 := math.Round(2.0 * (float64(AdjustPositiveNegative(int64(NumMod) - 3))) / float64(LongMod))
	res2 := math.Round(2.0 * -2 / float64(LongMod))

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

// Note: not 100% uniform over the error, used only for testing
func SufficientlyCloseLongMod(n int64) int64 {
	randomBit := int64(RandomBit())
	randomAbsError := EnsureLongMod(RandomInt64(NewRandomGenerator()))>>5

	randomError := (randomAbsError & -1 + randomBit) | (LongModNeg(randomAbsError) & -randomBit)

	return LongModAdd(n, randomError)
}

func SufficientlyCloseLongPolynomial(lp *LongPolynomial) *LongPolynomial {
	var coefficients [1024]int64

	for i := 0; i < len(coefficients); i++ {
		coefficients[i] = SufficientlyCloseLongMod(lp.Coefficients[i])
	}

	return &LongPolynomial { Coefficients: coefficients }
}

func TestReconciliate(t *testing.T) {
	cases := 2000

	for i := 0; i < cases; i++ {
		v := EnsureLongMod(RandomInt64(NewRandomGenerator()))
		w := SufficientlyCloseLongMod(v)

		crv := CrossRound(v)
		mrv := ModularRound(v)

		res := Reconciliate(w, crv)

		if res != mrv {
			t.Errorf(`
				Reconciliation failed
				v = %v
				w = %v
				crv = %v
				mrv = %v
				res = Reconciliate(%v, %v) = %v
				`,
				v,
				w,
				crv,
				mrv,
				w,
				crv,
				res,
			)
			return
		}
	}

}

func TestPolynomialReconciliate(t *testing.T) {
	cases := 2000

	for c := 0; c < cases; c++ {
		v := NewRandomLongPolynomial()
		w := SufficientlyCloseLongPolynomial(v)

		crv := v.CrossRound()

		mrv := v.ModularRound()

		res := w.Reconciliate(crv)

		if res.Coefficients != mrv.Coefficients {
			t.Errorf("polynomial reconciliation failed, crv == recon = %v\n", crv.Coefficients == res.Coefficients)

			var diffCount int

			for i := 0; i < len(res.Coefficients); i++ {
				if res.Coefficients[i] != mrv.Coefficients[i] {
					diffCount++
				}
			}

			t.Logf("diffCount = %v\n", diffCount)
			return
		}
	}
}
