package polynomial

const NumMod uint32 = (((1 << 31) - 1) << 1) + 1

const TwoInv uint32 = (1<<31)

const ModNegTwo uint32 = (NumMod - 2)

const LongMod int64 = (int64(1)<<33 - 2)

// Mod

func EnsureMod(a uint32) uint32 {
	var mask uint32 = 0xFFFFFFFF

	b := uint32(1 ^ BitwiseEq(int64(a), int64(mask)))

	return a & BitmaskFromBit(b^1)
}

func ModAdd(a, b uint32) uint32 {
	a = EnsureMod(a)
	b = EnsureMod(b)

	c := a + b

	c += uint32(BitwiseLt(int64(c), int64(a)))

	return EnsureMod(c)
}

func ModSubtract(a, b uint32) uint32 {
	return ModAdd(a, ModNeg(b))
}

func ModNeg(a uint32) uint32 {
	amod := EnsureMod(a)

	return EnsureMod(NumMod - amod)
}

func ModSub(a, b uint32) uint32 {
	return ModAdd(EnsureMod(a), ModNeg(b))
}

func ModMultiply(a, b uint32) uint32 {
	var c uint64

	a = EnsureMod(a)
	b = EnsureMod(b)

	c = uint64(a) * uint64(b)

	ans := ModAdd(uint32(c), uint32(c>>32))

	return ans
}


// LongMod

func EnsureLongMod(a int64) int64 {
	var mask int64 = 0x1FFFFFFFF

	a += LongMod

	q2 := ((a >> 33) << 1) & mask

	r := a & mask

	x1 := q2 + r

	q2 = ((x1 >> 33) << 1) & mask
	r = x1 & mask

	x1 = q2 + r

	// at this point x1 should be less than 2^33
	// so we just deduct 2 ^ 33 - 2

	// 33 bits all ones if x1 >= 2^33 - 2, otherwise zero
	firstOp := (int64(1) << 33) - 3

	modulo := (int64(1) << 33)

	compRes := BitwiseLt(firstOp, x1)

	flag := (modulo - compRes) & mask

	return x1 - (((1 << 33) - 2) & flag)
}

func AdjustPositiveNegative(n int64) int64 {
	// we want to leave the number as is if it's in the range [0, LongMod/2 - 1],
	// otherwise we want to shift back by LongMod (for example moving LongMod-1 to
	// -1) as it actually corresponds to -1 (goes back to 0 on adding 1)
	return n - (LongMod * (1 - BitwiseLt(n, (LongMod>>1))))
}

func LongModAdd(a, b int64) int64 {
	a = EnsureLongMod(a)
	b = EnsureLongMod(b)

	c := a + b

	return EnsureLongMod(c)
}

func LongModSubtract(a, b int64) int64 {
	return LongModAdd(a, LongModNeg(b))
}

func LongModMultiply(a, b int64) int64 {
	a = EnsureLongMod(a)
	b = EnsureLongMod(b)

	c := a * b

	return EnsureLongMod(c)
}

func LongModNeg(a int64) int64 {
	amod := EnsureLongMod(a)
	diff := LongMod - amod

	res := LongModAdd(diff, LongMod)

	return res
}

// the range down to q/4 is inclusive, so we set the value to floor(-q/4), we
// already know that q/4 is not an integer.
var ModularRoundLowerBound int64 = -(int64(1) << 31)

// the range up to q/4 is exclusive, so we set the value to floor(q/4), we
// already know that q/4 is not an integer.
var ModularRoundUpperBound int64 = (int64(1) << 31) - 1

func ModularRound(n int64) uint32 {
	nadj := AdjustPositiveNegative(n)
	res := (1 ^ BitwiseGtOrEqual(nadj, ModularRoundLowerBound)&BitwiseLt(nadj, ModularRoundUpperBound))

	return uint32(res & 1)
}

func CrossRound(n int64) uint32 {
	a := BitwiseLt(n, CrossRoundQ4)
	b := BitwiseLt(n, CrossRoundQ2)
	c := BitwiseLt(n, CrossRoundQ4+CrossRoundQ2)

	return 1 - uint32(a^b^c)
}

// here by q we mean (2^32 - 1)
// -q/4 = -(2^32 - 1)/4 = -2^30 + 0.25, so we round up to get the first element
// inside the set
var IZeroELowerBoundInclusive int64 = LongModNeg(1 << 30)

// round(q/2) + q/4 -1 = round(2 ^ 31 - 0.5) + (2 ^ 30 - 0.25) - 1 =
// 2^31 + 2^30 - 1.25, so we round up
// to get the first element outside the set = 2^31 + 2^30 - 1
var IZeroEUpperBoundExclusive int64 = (1 << 31) + (1 << 30) - 1

// -q/4 - floor(q/2) = -2^30 + 0.25 - floor(2^31 - 0.5) =
// -2^30 - 2^31 + 1 + 0.25, so we round up getting
// -2^30 - 2^31 + 1
var IOneELowerBoundInclusive int64 = LongModAdd(
	LongModAdd(LongModNeg(int64(1)<<31), LongModNeg(int64(1)<<30)),
	int64(1),
)

// q/4 - 1 = 2^30 - 0.25 - 1 = 2^30 - 1.25, so we round up getting 2^30 - 1
var IOneEUpperBoundExclusive int64 = (1 << 30) - 1

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
