package main

import "testing"

func TestNewBitField(t *testing.T) {
	bf := NewBitField(10, 10)
	if bf.width != 10 || bf.height != 10 {
		t.Errorf("Expected width=10, height=10, got width=%d, height=%d", bf.width, bf.height)
	}
	expectedLen := (100 + 63) / 64
	if len(bf.data) != expectedLen {
		t.Errorf("Expected data length %d, got %d", expectedLen, len(bf.data))
	}
}

func TestReadDefaultValue(t *testing.T) {
	bf := NewBitField(10, 10)
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if bf.Read(x, y) {
				t.Errorf("Expected false at (%d, %d), got true", x, y)
			}
		}
	}
}

func TestSetAndRead(t *testing.T) {
	bf := NewBitField(10, 10)

	bf.Set(0, 0, true)
	if !bf.Read(0, 0) {
		t.Errorf("Expected true at (0, 0), got false")
	}

	bf.Set(5, 5, true)
	if !bf.Read(5, 5) {
		t.Errorf("Expected true at (5, 5), got false")
	}

	bf.Set(9, 9, true)
	if !bf.Read(9, 9) {
		t.Errorf("Expected true at (9, 9), got false")
	}
}

func TestSetFalse(t *testing.T) {
	bf := NewBitField(10, 10)

	bf.Set(3, 3, true)
	if !bf.Read(3, 3) {
		t.Errorf("Expected true at (3, 3), got false")
	}

	bf.Set(3, 3, false)
	if bf.Read(3, 3) {
		t.Errorf("Expected false at (3, 3), got true")
	}
}

func TestIsolation(t *testing.T) {
	bf := NewBitField(10, 10)

	bf.Set(5, 5, true)

	if bf.Read(5, 4) || bf.Read(5, 6) || bf.Read(4, 5) || bf.Read(6, 5) {
		t.Errorf("Setting bit at (5, 5) affected neighboring bits")
	}
}

func TestLargeBitField(t *testing.T) {
	bf := NewBitField(100, 100)

	bf.Set(0, 0, true)
	bf.Set(99, 99, true)
	bf.Set(50, 50, true)

	if !bf.Read(0, 0) || !bf.Read(99, 99) || !bf.Read(50, 50) {
		t.Errorf("Large bitfield failed to set/read correctly")
	}

	if bf.Read(0, 1) || bf.Read(99, 98) || bf.Read(50, 51) {
		t.Errorf("Large bitfield has incorrect bits set")
	}
}

func TestAllPositions(t *testing.T) {
	bf := NewBitField(8, 8)

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			bf.Set(x, y, true)
		}
	}

	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			if !bf.Read(x, y) {
				t.Errorf("Expected true at (%d, %d), got false", x, y)
			}
		}
	}
}
