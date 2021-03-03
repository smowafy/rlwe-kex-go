package main

import(
//	"log"
)

// Operations will be in the ring Zq[X]/(X^1024 + 1) where Zq = Z / (2^32 - 1)Z
// in other words the polynomials are mod (X^1024 + 1) and the coefficients are
// mod (2^32 - 1)

const NumMod uint32 = (((1<<31) - 1)<<1) + 1

const LongMod int64 = (int64(1)<<33 - 2)

const CrossRoundQ4 int64 = (1<<31)
const CrossRoundQ2 int64 = int64(NumMod)

// TODO: change type to just [1024]uint32, no need for struct I guess?
type Polynomial struct {
	Coefficients [1024]uint32
}

func NewPolynomial() *Polynomial {
	return &Polynomial { }
}

type LongPolynomial struct {
	Coefficients [1024]int64
}

func NewLongPolynomial() *LongPolynomial {
	return &LongPolynomial { }
}

// TODO: remove if condition, side channel
func EnsureMod(a uint32) uint32 {
	if a == NumMod {
		return 0
	} else {
		return a
	}
}

func EnsureLongMod(a int64) int64 {
	var mask int64 = 0x1FFFFFFFF

	a += LongMod

	q2 := ((a>>33)<<1)&mask

	r := a & mask

	x1 := q2 + r

	q2 = ((x1>>33)<<1)&mask
	r = x1 & mask

	x1 = q2 + r

	// at this point x1 should be less than 2^33
	// so we just deduct 2 ^ 33 - 2

	// 33 bits all ones if x1 >= 2^33 - 2, otherwise zero
	firstOp := (int64(1)<<33)-3

	modulo := (int64(1)<<33)

	compRes := BitwiseLt(firstOp, x1)

	flag := (modulo - compRes)&mask

	return x1 - (((1<<33) - 2) & flag)
}

// x ^ ((x ^ y) | (x - y) ^ x)
// taken from the original repo:
// 	https://github.com/dstebila/rlwekex/blob/master/rlwe.c
func BitwiseLt(a, b int64) int64 {
	return (a ^ ((a ^ b) | (a - b) ^ a))>>63 & 1
}

// also taken from the repo
func BitwiseEq(a, b int64) int64 {
	return ((a - b) | (b - a)) >> 63;
}

func BitwiseLtOrEqual(a, b int64) int64 {
	return 1 ^ BitwiseLt(b, a)
}

func ModAdd(a, b uint32) uint32 {
	a = EnsureMod(a)
	b = EnsureMod(b)

	c := a + b

	// TODO: remove if condition, side channel
	if c < a {
		c++
	}

	return EnsureMod(c)
}

func ModMultiply(a, b uint32) uint32 {
	var c uint64

	a = EnsureMod(a)
	b = EnsureMod(b)

	c = uint64(a) * uint64(b)

	return ModAdd(uint32(c), uint32(c>>32))
}

func ModNeg(a uint32) uint32 {
	amod := EnsureMod(a)
	diff := NumMod - amod

	res := ModAdd(diff, NumMod)
//	return ModAdd(NumMod - EnsureMod(a), NumMod)

	return res
}

func ModSub(a, b uint32) uint32 {
	return ModAdd(EnsureMod(a), ModNeg(b))
}

func (p *Polynomial) Add(other *Polynomial) *Polynomial {
	for i := 0; i < 1024; i++ {
		p.Coefficients[i] = ModAdd(p.Coefficients[i], other.Coefficients[i])
	}

	return p
}

// TODO: Make more efficient: FFT and shit
func (p *Polynomial) Multiply(other *Polynomial) *Polynomial {
	var res uint32

	pcop := Polynomial {}

	for i := 0; i < 1024; i++ {
		for j := 0; j < 1024; j++ {
			res = ModMultiply(p.Coefficients[i], other.Coefficients[j])


			// TODO: remove if condition, side channel potential
			if i + j > 1024 {
				res = ModNeg(res)
			}

			// mod 1024
			idx := (i + j) & (1<<10 - 1)

			pcop.Coefficients[idx] += res

		}
	}

	for i := 0; i < 1024; i++ {
		p.Coefficients[i] = pcop.Coefficients[i]
	}

	return p
}

func (p *Polynomial) ErrorDouble(rg RandomGenerator) *LongPolynomial {
	res := NewLongPolynomial()

	doubleError := RandomBit() + (-1 * RandomBit())

	for i := 0; i < len(p.Coefficients); i++ {
		res.Coefficients[i] = EnsureLongMod(int64(p.Coefficients[i]<<1) + int64(doubleError))
	}

	return res
}

func (lp *LongPolynomial) ModularRound() *Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < 1024; i++ {
		res := BitwiseLt(lp.Coefficients[i], (LongMod>>1))

		coefficients[i] = 1 - uint32(res)
	}

	return &Polynomial{ Coefficients: coefficients }
}

func (lp *LongPolynomial) CrossRound() *Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < 1024; i++ {
		a := BitwiseLt(lp.Coefficients[i], CrossRoundQ4)
		b := BitwiseLt(lp.Coefficients[i], CrossRoundQ2)
		c := BitwiseLt(lp.Coefficients[i], CrossRoundQ4 + CrossRoundQ2)

		coefficients[i] = 1 - uint32(a ^ b ^ c)
	}

	return &Polynomial { Coefficients: coefficients }
}

func (p *Polynomial) Reconciliate(b uint32) *Polynomial {
	// hard-code the bounds of I0+E and I1+E
	// mask according to bit to get Ib+E
	// not(lower_bound <= w < upper_bound)
	return p
}
