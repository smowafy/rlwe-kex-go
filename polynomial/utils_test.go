package polynomial

func TestBitmaskFromBit(t *testing.T) {
	sample := 0

	if BitmaskFromBit(uint32(sample)) != 0 {
		t.Errorf("Failed\n")
	}

	sample = 1

	if BitmaskFromBit(uint32(sample)) != 0xFFFFFFFF {
		t.Errorf("Failed\n")
	}
}
