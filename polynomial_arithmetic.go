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

func RandomUInt32(rg RandomGenerator) uint32 {
	var ans uint32

	b := make([]byte, 4)
	_, err := rg.Read(b)

	if err != nil {
		panic(err)
	}

	for j := 0; j < 4; j++ {
		ans |= uint32(b[j])<<(j<<3)
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
		ans |= int64(b[j])<<(j<<3)
	}

	return ans
}

func NewPolynomial() Polynomial {
	return Polynomial { }
}

func NewSmallPolynomial() Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < len(coefficients); i++ {
		coefficients[i] = EnsureMod(uint32(GaussianOverZ()))
	}

	return Polynomial { Coefficients: coefficients }
}


func NewRandomPolynomial() Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < len(coefficients); i++ {
		coefficients[i] = EnsureMod(RandomUInt32(NewRandomGenerator()))
	}

	return Polynomial { Coefficients: coefficients }
}

type LongPolynomial struct {
	Coefficients [1024]int64
}

func NewRandomLongPolynomial() LongPolynomial {
	var coefficients [1024]int64

	for i := 0; i < len(coefficients); i++ {
		coefficients[i] = EnsureLongMod(RandomInt64(NewRandomGenerator()))
	}

	return LongPolynomial { Coefficients: coefficients }
}

func NewLongPolynomialFromPolynomial(p Polynomial) LongPolynomial {
	var coefficients [1024]int64

	for i := 0; i < len(coefficients); i++ {
		coefficients[i] = int64(p.Coefficients[i])
	}

	return LongPolynomial { Coefficients: coefficients }
}


// produces 0 if given 0 (NumMod + 1)
// produces 2^32 - 1 (NumMod) if given 1
func BitmaskFromBit(b uint32) uint32 {
	return NumMod + ((b & 1) ^ 1)
}

func EnsureMod(a uint32) uint32 {
	var mask uint32 = 0xFFFFFFFF

	b := uint32(1 ^ BitwiseEq(int64(a), int64(mask)))

	return a & BitmaskFromBit(b ^ 1)
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

func AdjustPositiveNegative(n int64) int64 {
	// we want to leave the number as is if it's in the range [0, LongMod/2 - 1],
	// otherwise we want to shift back by LongMod (for example moving LongMod-1 to
	// -1) as it actually corresponds to -1 (goes back to 0 on adding 1)
	return n - (LongMod * (1 - BitwiseLt(n, (LongMod>>1))))
}

// 1 if a < b, otherwise 0
// a ^ ((a ^ b) | (a - b) ^ a)
// taken from the original repo:
// 	https://github.com/dstebila/rlwekex/blob/master/rlwe.c
func BitwiseLt(a, b int64) int64 {
	return (a ^ ((a ^ b) | (a - b) ^ a))>>63 & 1
}

// also taken from the repo
func BitwiseEq(a, b int64) int64 {
	return ((a - b) | (b - a)) >> 63;
}

func BitwiseGtOrEqual(a, b int64) int64 {
	return BitwiseLt(b, a)
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

func LongModAdd(a, b int64) int64 {
	a = EnsureLongMod(a)
	b = EnsureLongMod(b)

	c := a + b

	// TODO: remove if condition, side channel
	return EnsureLongMod(c)
}

func LongModSubtract(a, b int64) int64 {
	return LongModAdd(a, LongModNeg(b))
}

func ModMultiply(a, b uint32) uint32 {
	var c uint64

	a = EnsureMod(a)
	b = EnsureMod(b)

	c = uint64(a) * uint64(b)

	return ModAdd(uint32(c), uint32(c>>32))
}

func LongModMultiply(a, b int64) int64 {
	a = EnsureLongMod(a)
	b = EnsureLongMod(b)

	c := a * b

	return EnsureLongMod(c)
}

func ModNeg(a uint32) uint32 {
	amod := EnsureMod(a)
	diff := NumMod - amod

	res := ModAdd(diff, NumMod)
//	return ModAdd(NumMod - EnsureMod(a), NumMod)

	return res
}

func LongModNeg(a int64) int64 {
	amod := EnsureLongMod(a)
	diff := LongMod - amod

	res := LongModAdd(diff, LongMod)

	return res
}

func ModSub(a, b uint32) uint32 {
	return ModAdd(EnsureMod(a), ModNeg(b))
}

func (p Polynomial) Add(other Polynomial) Polynomial {
	res := Polynomial{}

	for i := 0; i < 1024; i++ {
		res.Coefficients[i] = ModAdd(p.Coefficients[i], other.Coefficients[i])
	}

	return res
}

// TODO: Make more efficient: FFT and shit
func (p Polynomial) Multiply(other Polynomial) Polynomial {
	var res uint32

	pcop := Polynomial {}

	for i := 0; i < 1024; i++ {
		for j := 0; j < 1024; j++ {
			res = ModMultiply(p.Coefficients[i], other.Coefficients[j])


			// TODO: remove if condition, side channel potential
			if i + j >= 1024 {
				res = ModNeg(res)
			}

			// mod 1024
			idx := (i + j) & ((1<<10) - 1)

			pcop.Coefficients[idx] = ModAdd(pcop.Coefficients[idx], res)

		}
	}

	return pcop
}

// Quite an assumption here: `other` is a small polynomial (generated from the
// GaussianOverZ function with very low probability to have absolute value
// greater than 51
func (p LongPolynomial) Multiply(other LongPolynomial) LongPolynomial {
	var res int64

	pcop := LongPolynomial{}

	for i := 0; i < 1024; i++ {
		for j := 0; j < 1024; j++ {
			res = LongModMultiply(p.Coefficients[i], other.Coefficients[j])


			// TODO: remove if condition, side channel potential
			if i + j >= 1024 {
				res = LongModNeg(res)
			}

			// mod 1024
			idx := (i + j) & (1<<10 - 1)

			pcop.Coefficients[idx] = LongModAdd(pcop.Coefficients[idx], res)

		}
	}

	return pcop
}

func (p Polynomial) Double() LongPolynomial {
	res := LongPolynomial{}

	for i := 0; i < len(p.Coefficients); i++ {
		res.Coefficients[i] = EnsureLongMod(int64(p.Coefficients[i])<<1)
	}

	return res
}

func (p Polynomial) ErrorDouble(rg RandomGenerator) LongPolynomial {
	res := LongPolynomial{}

	doubleError := RandomBit() + (-1 * RandomBit())

	for i := 0; i < len(p.Coefficients); i++ {
		res.Coefficients[i] = EnsureLongMod((int64(p.Coefficients[i])<<1) + int64(doubleError))
	}

	return res
}

// the range down to q/4 is inclusive, so we set the value to floor(-q/4), we
// already know that q/4 is not an integer.
var ModularRoundLowerBound int64 = -(int64(1)<<31)

// the range up to q/4 is exclusive, so we set the value to floor(q/4), we
// already know that q/4 is not an integer.
var ModularRoundUpperBound int64 = (int64(1)<<31) - 1

func ModularRound(n int64) uint32 {
	nadj := AdjustPositiveNegative(n)
	res := (1 ^ BitwiseGtOrEqual(nadj, ModularRoundLowerBound) & BitwiseLt(nadj, ModularRoundUpperBound))

	return uint32(res & 1)
}

func (lp LongPolynomial) ModularRound() Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < 1024; i++ {
		coefficients[i] = ModularRound(lp.Coefficients[i])
	}

	return Polynomial{ Coefficients: coefficients }
}

func CrossRound(n int64) uint32 {
	a := BitwiseLt(n, CrossRoundQ4)
	b := BitwiseLt(n, CrossRoundQ2)
	c := BitwiseLt(n, CrossRoundQ4 + CrossRoundQ2)
	return 1 - uint32(a ^ b ^ c)
}

func (lp LongPolynomial) CrossRound() Polynomial {
	var coefficients [1024]uint32

	for i := 0; i < 1024; i++ {
		coefficients[i] = CrossRound(lp.Coefficients[i])
	}

	return Polynomial { Coefficients: coefficients }
}

// here by q we mean (2^32 - 1)
// -q/4 = -(2^32 - 1)/4 = -2^30 + 0.25, so we round up to get the first element
// inside the set
var IZeroELowerBoundInclusive int64 = LongModNeg(1<<30)

// round(q/2) + q/4 -1 = round(2 ^ 31 - 0.5) + (2 ^ 30 - 0.25) - 1 =
// 2^31 + 2^30 - 1.25, so we round up
// to get the first element outside the set = 2^31 + 2^30 - 1
var IZeroEUpperBoundExclusive int64 = (1<<31) + (1<<30) - 1

// -q/4 - floor(q/2) = -2^30 + 0.25 - floor(2^31 - 0.5) =
// -2^30 - 2^31 + 1 + 0.25, so we round up getting
// -2^30 - 2^31 + 1
var IOneELowerBoundInclusive int64 = LongModAdd(
	LongModAdd(LongModNeg(int64(1)<<31), LongModNeg(int64(1)<<30)),
	int64(1),
)

// q/4 - 1 = 2^30 - 0.25 - 1 = 2^30 - 1.25, so we round up getting 2^30 - 1
var IOneEUpperBoundExclusive int64 = (1<<30) - 1

func Reconciliate(w int64, b uint32) uint32 {
	// hard-code the bounds of I0+E and I1+E
	// mask according to bit to get Ib+E
	// not(lower_bound <= w < upper_bound)

	// strongly assuming that b is a bit here, to obtain the masks using 0 and -1

	IBLowerBound := ((-1 + int64(b)) & IZeroELowerBoundInclusive) | ((-int64(b)) & IOneELowerBoundInclusive)
	IBUpperBound := ((-1 + int64(b)) & IZeroEUpperBoundExclusive) | ((-int64(b)) & IOneEUpperBoundExclusive)

	inRange := BitwiseGtOrEqual(AdjustPositiveNegative(w), AdjustPositiveNegative(IBLowerBound)) &
		BitwiseLt(AdjustPositiveNegative(w), AdjustPositiveNegative(IBUpperBound))

	res := (1 ^ inRange) & 1

	return uint32(res)
//	return uint32((BitwiseGtOrEqual(IBLowerBound, w) & BitwiseLt(w, IBUpperBound)))
}

func (w LongPolynomial) Reconciliate(b Polynomial) Polynomial {
	res := Polynomial{}

	for i := 0; i < len(w.Coefficients); i++ {
		res.Coefficients[i] = Reconciliate(w.Coefficients[i], b.Coefficients[i])
	}

	return res
}
