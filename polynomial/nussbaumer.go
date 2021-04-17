package polynomial

import(
	"github.com/smowafy/rlwe-kex-go/utils"
)

// correct only for power of 2 lengths
func lookupShift(a []uint32, shift int, i int, log2a int) uint32 {
	// for mod
	mask := len(a) - 1

	idx := shift + i

	q := (idx >> log2a)
	r := (idx & mask)

	sign := uint32(q & 1)

	ans := (BitmaskFromBit(sign) & ModNeg(a[r])) | (BitmaskFromBit(^sign) & a[r])

	return ans
}

func NussbaumerPolynomial(inp []uint32) [][]uint32 {
	log2len := utils.Log2(len(inp))

	// closest power of two to the square root??
	l1 := (1<<(log2len>>1))
	l2 := len(inp)>>(log2len>>1)

	res := make([][]uint32, l1)

	for i := 0; i < l1; i++ {
		res[i] = make([]uint32, l2)

		for j := 0; j < l2; j++ {
			res[i][j] = inp[j * l1 + i]
		}
	}

	return res
}

func NussbaumerDoublePolynomial(inp []uint32) [][]uint32 {
	log2len := utils.Log2(len(inp))

	// closest power of two to the square root??
	l1 := (1<<(log2len>>1))
	l2 := len(inp)>>(log2len>>1)

	res := make([][]uint32, (l1<<1))

	for i := 0; i < l1; i++ {
		res[i] = make([]uint32, l2)
		res[i+l1] = make([]uint32, l2)

		for j := 0; j < l2; j++ {
			res[i][j] = inp[j * l1 + i]
		}
	}

	return res
}

func NussbaumerInversePolynomial(inp [][]uint32) []uint32 {
	l1 := len(inp)
	l2 := len(inp[0])
	res := make([]uint32, l1*l2)

	for i := range inp {
		for j := range inp[i] {
			res[j * l1 + i] = inp[i][j]
		}
	}

	return res
}

func MostSignificantBitIndex(a int) int {
	mostSignificantBitIndex := 0

	mask := 1

	bit := 0

	for i := 0; i < 32; i++ {
		bit = (a & mask)>>i

		mostSignificantBitIndex = mostSignificantBitIndex - (bit * mostSignificantBitIndex) + (bit * i)

		mask<<=1
	}

	return mostSignificantBitIndex
}

// TODO: optimize and generate the numbers (0 -> 2^n) instead of calculating reversal
func BitReverse(a int, mostSignificantBitIndex int) int {
	ans := 0

	for i := mostSignificantBitIndex; i >= 0; i-- {
		ans |= (((a>>i)&1)<<(mostSignificantBitIndex - i))
	}

	return ans
}

func NussbaumerIterativePolynomial(a []uint32) [][]uint32 {
	msbi := (MostSignificantBitIndex(len(a)) - 1)

	ans := make([][]uint32, len(a))

	for i := 0; i < len(a); i++ {
		ans[i] = make([]uint32, (len(a)>>1))

		ans[i][0] = a[BitReverse(i, msbi)]
	}

	return ans
}

func NussbaumerIterativeReorderPolynomial(a [][]uint32) [][]uint32 {
	msbi := (MostSignificantBitIndex(len(a)) - 1)

	ans := make([][]uint32, len(a))

	for i := 0; i < len(a); i++ {
		ans[i] = make([]uint32, (len(a)>>1))

		ans[i] = a[BitReverse(i, msbi)]
	}

	return ans
}

func NussbaumerIterativeTransform(b [][]uint32) [][]uint32 {
	n := len(b)
	a := b

	log2a := utils.Log2(n)

	a0, a1 := make([]uint32, len(a[0])), make([]uint32, len(a[0]))

	innerJmp := len(a[0])

	for l := 2; l <= n; l<<=1 {
		for i := 0; i < n; i+=l {
			for j := 0; j < (l>>1); j++ {
				shiftVal := -j * innerJmp

				copy(a0, a[i+j])
				copy(a1, a[i+j+(l>>1)])

				for k := range a[i+j] {
					lookup := lookupShift(a1, shiftVal, k, log2a)
					a[i+j][k] = ModAdd(a0[k], lookup)

					a[i+j+(l>>1)][k] = ModAdd(a0[k], ModNeg(lookup))
				}
			}
		}
		innerJmp>>=1
	}

	return a
}

func NussbaumerIterativeInverseTransform(b [][]uint32) [][]uint32 {
	n := len(b)
	a := b

	log2a := utils.Log2(len(a[0]))

	a0, a1 := make([]uint32, len(a[0])), make([]uint32, len(a[0]))

	innerJmp := len(a[0])
	for l := 2; l <= n; l<<=1 {
		for i := 0; i < n; i+=l {
			for j := 0; j < (l>>1); j++ {
				shiftVal := j * innerJmp

				copy(a0, a[i+j])
				copy(a1, a[i+j+(l>>1)])

				for k := range a[i+j] {
					lookup := lookupShift(a1, shiftVal, k, log2a)
					a[i+j][k] = ModMultiply(ModAdd(a0[k], lookup), TwoInv)

					a[i+j+(l>>1)][k] = ModMultiply(ModAdd(a0[k], ModNeg(lookup)), TwoInv)
				}
			}
		}
		innerJmp>>=1
	}

	return a
}


// Recursive
func NussbaumerTransform(ahat [][]uint32) [][]uint32 {

	l2 := len(ahat)

	if l2 == 1 {

		return ahat
	}

	ahat0 := make([][]uint32, l2>>1)
	ahat1 := make([][]uint32, l2>>1)

	for i := 0; i < (l2 >> 1); i++ {
		ahat0[i] = ahat[i<<1]
		ahat1[i] = ahat[(i<<1)+1]
	}

	a0 := NussbaumerTransform(ahat0)
	a1 := NussbaumerTransform(ahat1)

	ans := make([][]uint32, l2)

	for i := range ans {
		ans[i] = make([]uint32, (len(a0[0])))
	}

	for i := 0; (i << 1) < l2; i++ {

		for j := 0; j < len(a0[i]); j++ {

			// Note: it's a division operation that's technically not constant time,
			// however it doesn't depend on the input but on the slice lengths, which
			// are pretty much constant for the implementation, so supposedly no risk
			// here (even though the iterative version is actually used which does not
			// even include division).
			shiftVal := -i * len(a0[i]) / len(a0)
			ans[i][j] = ModAdd(a0[i][j], lookupShift(a1[i], shiftVal, j, utils.Log2(len(a1[i]))))

			ans[i+(len(ans)>>1)][j] = ModAdd(a0[i][j], ModNeg(lookupShift(a1[i], shiftVal, j, utils.Log2(len(a1[i])))))
		}

	}

	return ans
}

func NussbaumerInverseTransform(ahat [][]uint32) [][]uint32 {
	l2 := len(ahat)

	if l2 == 1 {

		return ahat
	}

	ahat0 := make([][]uint32, l2>>1)
	ahat1 := make([][]uint32, l2>>1)

	for i := 0; i < (l2 >> 1); i++ {
		ahat0[i] = ahat[i<<1]
		ahat1[i] = ahat[(i<<1)+1]
	}

	a0 := NussbaumerInverseTransform(ahat0)
	a1 := NussbaumerInverseTransform(ahat1)

	ans := make([][]uint32, l2)

	for i := range ans {
		ans[i] = make([]uint32, (len(a0[0])))
	}

	for i := 0; (i << 1) < l2; i++ {

		for j := 0; j < len(a0[i]); j++ {

			// Note: it's a division operation that's technically not constant time,
			// however it doesn't depend on the input but on the slice lengths, which
			// are pretty much constant for the implementation, so supposedly no risk
			// here (even though the iterative version is actually used which does not
			// even include division).
			shiftVal := i * len(a0[i]) / len(a0)
			ans[i][j] = ModMultiply(ModAdd(a0[i][j], lookupShift(a1[i], shiftVal, j, utils.Log2(len(a1[i])))), TwoInv)

			ans[i+(len(ans)>>1)][j] = ModMultiply(ModAdd(a0[i][j], ModNeg(lookupShift(a1[i], shiftVal, j, utils.Log2(len(a1[i]))))), TwoInv)
		}

	}

	return ans
}


// assuming a and b are of the same length and it is a power of two
func NaiveMultiply(a []uint32, b []uint32) []uint32 {
	ans := make([]uint32, len(a))

	var res uint32

	for i, ai := range a {
		for j, bj := range b {
			res = ModMultiply(ai, bj)

			if i+j >= len(a) {
				res = ModNeg(res)
			}

			idx := (i+j) & (len(a) - 1)

			ans[idx] = ModAdd(ans[idx], res)
		}
	}

	return ans
}

func NussbaumerRecursiveMultiply(a []uint32, b []uint32) []uint32 {
	if len(a) <= 32 {
		return NaiveMultiply(a, b)
	}

	az := NussbaumerDoublePolynomial(a)
	bz := NussbaumerDoublePolynomial(b)

	azhat := NussbaumerTransform(az)
	bzhat := NussbaumerTransform(bz)

	czhat := make([][]uint32, len(azhat))

	for i := range azhat {
		czhat[i] = NussbaumerRecursiveMultiply(azhat[i], bzhat[i])
	}

	cz := NussbaumerInverseTransform(czhat)

	cex := make([]uint32, (len(a)<<1))

	l1 := len(cz)>>1

	for i := range cz {
		for j := range cz[i] {
			cex[j * l1 + i] = ModAdd(cex[j * l1 + i], cz[i][j])
		}
	}

	c := make([]uint32, len(a))

	copy(c, cex)

	for i := len(c); i < len(cex); i++ {
		c[i - len(c)] = ModSubtract(c[i - len(c)], cex[i])
	}

	return c
}

func NussbaumerIterativeMultiply(a []uint32, b []uint32) []uint32 {
	if len(a) <= 32 {
		return NaiveMultiply(a, b)
	}

	az := NussbaumerDoublePolynomial(a)
	bz := NussbaumerDoublePolynomial(b)

	azForFft := NussbaumerIterativeReorderPolynomial(az)
	bzForFft := NussbaumerIterativeReorderPolynomial(bz)

	azhat := NussbaumerIterativeTransform(azForFft)
	bzhat := NussbaumerIterativeTransform(bzForFft)

	czhat := make([][]uint32, len(azhat))

	for i := range azhat {
		czhat[i] = NussbaumerIterativeMultiply(azhat[i], bzhat[i])
	}

	czhatForFft := NussbaumerIterativeReorderPolynomial(czhat)

	cz := NussbaumerIterativeInverseTransform(czhatForFft)

	cex := make([]uint32, (len(a)<<1))

	l1 := len(cz)>>1

	for i := range cz {
		for j := range cz[i] {
			cex[j * l1 + i] = ModAdd(cex[j * l1 + i], cz[i][j])
		}
	}

	c := make([]uint32, len(a))

	copy(c, cex)

	for i := len(c); i < len(cex); i++ {
		c[i - len(c)] = ModSubtract(c[i - len(c)], cex[i])
	}

	return c
}
