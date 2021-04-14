package polynomial

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

	for i = 0; i < 10000000; i++ {
		expectedRes := uint32(int64(-i) % (int64(1)<<33 - 2))
		res := EnsureMod(-i)
		if res != EnsureMod(expectedRes) {
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

	prev := -1

	t.Logf("uint32(-1) = %v\n", uint32(prev))

	sign := uint32(prev) + NumMod

	if EnsureMod(uint32(sign)) != NumMod-1 {
		t.Errorf("mod failed, expected %v but got %v\n", NumMod-1, EnsureMod(uint32(sign)))
	}
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

	for i = 0; i < (1 << 60); i += rand.Int63n(1<<40) + (1 << 30) {
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
		c = rand.Int63n(LongMod >> 2)

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
		a, b = rand.Int63n((1 << 62)), rand.Int63n((1 << 62))

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
			t.Errorf("mod multiplication failed, expected %v but got %v\n", expectedRes, res)
		}
	}
}

func TestLongModMultiply(t *testing.T) {
	var a, b, res int64
	var expectedRes int64

	const cases int = 1000

	for c := 0; c < cases; c++ {
		// testing with one of the numbers slightly smaller not to overflow in
		// multiplication
		a, b = EnsureLongMod(rand.Int63()), (rand.Int63() & 0xFFFFFFF)

		expectedRes = (((a * b) % LongMod) + LongMod) % LongMod

		res = LongModMultiply(a, b)

		if EnsureLongMod(expectedRes) != res {
			t.Errorf("long mod multiply failed, expected %v but got %v\n", expectedRes, res)
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
	for a := (int64(1) << 32); a <= (int64(1) << 34); a += (rand.Int63n(int64(1) << 20)) {
		expected := (-a + 8*LongMod) % LongMod
		actual := LongModNeg(a)

		if expected != actual {
			t.Errorf("mod negation modulo %v failed for %v, expected %v but got %v\n", LongMod, a, expected, actual)

			return
		}
	}
}
