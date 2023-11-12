package calculations

import (
	"math"
	"testing"
)

// TestAdd1 testet die Funktion Add1.
func TestAdd1(t *testing.T) {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			result := Add1(x, y)
			if result != x+y {
				t.Errorf("Add1(%d,%d) == %d != %d", x, y, result, x+y)
			}
		}
	}
}

// TestAdd2 testet die Funktion Add2.
func TestAdd2(t *testing.T) {
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			result := Add2(x, y)
			if result != x+y {
				t.Errorf("Add2(%d,%d) == %d != %d", x, y, result, x+y)
			}
		}
	}
}

// TestPower testet die Funktion Power.
func TestPower(t *testing.T) {
	for x := 0; x < 10; x++ {
		for n := 0; n < 10; n++ {
			result := Power(x, n)
			expected := int(math.Pow(float64(x), float64(n)))
			if result != expected {
				t.Errorf("Power(%d,%d) == %d != %d", x, n, result, expected)
			}
		}
	}
}
