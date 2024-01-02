package game

import (
	"testing"
)

func TestNewPieceFromSliceEmpty(t *testing.T) {
	var emptyBuf [][]bool
	_, err := NewPieceFromSlice(emptyBuf)
	if err == nil {
		t.Errorf("expected error from creating an empty piece")
	}
}

func TestNewPieceFromString(t *testing.T) {
	z4Duplicate, err := NewPieceFromString("OXX|XXO")
	if err != nil {
		t.Errorf("got error from creating new piece: %v", err)
	}
	if !z4Duplicate.Equals(Z4) {
		t.Errorf("Z4 from string does not match actual Z4 piece: %v", err)
	}
}

func TestRotate90DegreesRight(t *testing.T) {
	rotatedZ4 := Z4.Rotate90DegreesRight()
	actual := rotatedZ4.String()
	expected := "XO|XX|OX"
	if actual != expected {
		t.Errorf("expected rotated piece is %s, got %s", expected, actual)
	}
}

func TestString(t *testing.T) {
	actual := L4.String()
	expected := "XOO|XXX"
	if actual != expected {
		t.Errorf("expected piece string is %s, got %s", expected, actual)
	}
}

func TestEquals(t *testing.T) {
	rotatedU := U
	for i := 0; i < 4; i++ {
		if !rotatedU.Equals(U) {
			t.Errorf("expected equality to hold despite rotation")
		}
		rotatedU = rotatedU.Rotate90DegreesRight()
	}
}
