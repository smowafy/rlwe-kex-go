package main

// Operations will be in the ring Zq[X]/(X^1024 + 1) where Zq = Z / (2^32 - 1)Z
// in other words the polynomials are mod (X^1024 + 1) and the coefficients are
// mod (2^32 - 1)

const NumMod uint32 = (((1<<31) - 1)<<1) + 1

// TODO: change type to just [1024]uint32, no need for struct I guess?
type Polynomial struct {
	Coefficients [1024]uint32
}

func NewPolynomial() *Polynomial {
	return &Polynomial { }
}

// TODO: remove if condition, side channel
func EnsureMod(a uint32) uint32 {
	if a == NumMod {
		return 0
	} else {
		return a
	}
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
