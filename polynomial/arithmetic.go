package polynomial

import (
	"github.com/smowafy/rlwe-kex-go/gaussian"
)

// Operations will be in the ring Zq[X]/(X^1024 + 1) where Zq = Z / (2^32 - 1)Z
// in other words the polynomials are mod (X^1024 + 1) and the coefficients are
// mod (2^32 - 1)

type Polynomial []uint32

const CrossRoundQ4 int64 = (1 << 31)
const CrossRoundQ2 int64 = int64(NumMod)

func NewPolynomial() Polynomial {
	return make(Polynomial, 1024)
}

func NewSmallPolynomial() Polynomial {
	p := make(Polynomial, 1024)

	for i := range p {
		p[i] = EnsureMod(uint32(gaussian.GaussianOverZ()))
	}

	return p
}

func NewRandomPolynomial() Polynomial {
	p := make(Polynomial, 1024)

	for i := range p {
		p[i] = EnsureMod(RandomUInt32(gaussian.NewRandomGenerator()))
	}

	return p
}

type LongPolynomial []int64

func NewRandomLongPolynomial() LongPolynomial {
	p := make(LongPolynomial, 1024)

	for i := range p {
		p[i] = EnsureLongMod(RandomInt64(gaussian.NewRandomGenerator()))
	}

	return p
}

func NewLongPolynomialFromPolynomial(p Polynomial) LongPolynomial {
	lp := make(LongPolynomial, 1024)

	for i := range p {
		lp[i] = int64(p[i])
	}

	return lp
}


func (p Polynomial) Add(other Polynomial) Polynomial {
	res := make(Polynomial, len(p))

	for i := range p {
		res[i] = ModAdd(p[i], other[i])
	}

	return res
}


// assuming they are of the same length
func (p Polynomial) Multiply(other Polynomial) Polynomial {
	return NussbaumerIterativeMultiply(p, other)
}

func (p Polynomial) Double() LongPolynomial {
	res := make(LongPolynomial, len(p))

	for i := range p {
		res[i] = EnsureLongMod(int64(p[i]) << 1)
	}

	return res
}

func (p Polynomial) ErrorDouble(rg gaussian.RandomGenerator) LongPolynomial {
	res := make(LongPolynomial, len(p))

	doubleError := gaussian.RandomBit() + (-1 * gaussian.RandomBit())

	for i := range p {
		res[i] = EnsureLongMod((int64(p[i]) << 1) + int64(doubleError))
	}

	return res
}


func (lp LongPolynomial) ModularRound() Polynomial {
	p := make(Polynomial, len(lp))

	for i := range p {
		p[i] = ModularRound(lp[i])
	}

	return p
}

func (lp LongPolynomial) CrossRound() Polynomial {
	p := make(Polynomial, len(lp))

	for i := range p {
		p[i] = CrossRound(lp[i])
	}

	return p
}


func (w LongPolynomial) Reconciliate(b Polynomial) Polynomial {
	res := make(Polynomial, len(b))

	for i := range res {
		res[i] = Reconciliate(w[i], b[i])
	}

	return res
}
