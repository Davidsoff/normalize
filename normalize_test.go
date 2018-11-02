package normalize

import (
	"testing"
)

func TestMin(t *testing.T) {
	cases := []struct {
		in     []float64
		wanted float64
	}{
		{[]float64{1.0, 1.1, 1.3, 4, 20}, 1.0},
		{[]float64{-1.0, -1.1, -1.3, -4, -20}, -20},
		{[]float64{-1.0, 1.1, 1.3, 4, 20}, -1.0},
	}
	for _, c := range cases {
		got := Min(c.in)
		if got != c.wanted {
			t.Errorf("Min(%v) == %e, want %e", c.in, got, c.wanted)
		}
	}
}

func TestMax(t *testing.T) {
	cases := []struct {
		in     []float64
		wanted float64
	}{
		{[]float64{1.0, 1.1, 1.3, 4, 20}, 20},
		{[]float64{-1.0, -1.1, -1.3, -4, -20}, -1.0},
		{[]float64{-1.0, 1.1, 1.3, 4, 20}, 20},
	}
	for _, c := range cases {
		got := Max(c.in)
		if got != c.wanted {
			t.Errorf("Max(%v) == %e, want %e", c.in, got, c.wanted)
		}
	}
}

func TestNormalize(t *testing.T) {
	cases := []struct {
		in     []float64
		wanted []float64
	}{
		{[]float64{0, 1, 1.5, 2, 4, 10}, []float64{0, 0.1, 0.15, 0.2, 0.4, 1}},
		{[]float64{11, 11, 11, 11, 11}, []float64{1, 1, 1, 1, 1}},
	}
	for _, c := range cases {
		got := Normalize(c.in)
		if !testEq(got, c.wanted) {
			t.Errorf("Normalize(%v) == %e, want %e", c.in, got, c.wanted)
		}
	}
}

func testEq(a, b []float64) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
