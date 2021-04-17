package utils

import(
	"testing"
	"math/rand"
)

func TestBitwiseLt(t *testing.T) {
	cases := 2000000

	for ; cases > 0; cases-- {
		a := rand.Int63()
		b := rand.Int63()

		var expected int64

		if a < b {
			expected = 1
		} else {
			expected = 0
		}

		res := BitwiseLt(a, b)

		if res != expected {
			t.Errorf("bitwise LT failed for %v < %v\nexpected %v\ngot%v\n", a, b, expected, res)
			return
		}
	}

	a := int64(-2147483648)
	b := int64(1073741822)

	var expected int64

	if a < b {
		expected = 1
	} else {
		expected = 0
	}


	res := BitwiseGtOrEqual(b, a)

	if res != expected {
		t.Errorf("bitwise LT failed for %v < %v\nexpected %v\ngot%v\n", a, b, expected, res)
		return
	}

	a = int64(6129484611666145821)
	b = int64(6129484611666145821)

	if a < b {
		expected = 1
	} else {
		expected = 0
	}


	res = BitwiseGtOrEqual(b, a)

	if res != expected {
		t.Errorf("bitwise LT failed for %v < %v\nexpected %v\ngot%v\n", a, b, expected, res)
		return
	}
}
