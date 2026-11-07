package wtconv

import (
	"testing"

	"github.com/unixlinuxgeek/randbtw"
)

func TestLbToKg(t *testing.T) {
	n := randbtw.Generate(1, 9999)

	t.Logf("Generate-Random-Integer: %v\n", n)
	correct := Kilogram(float64(n) / 2.2046226218) // correct formula
	needCheck := LbToKg(Pound(n))

	if needCheck != correct {
		t.Errorf("\"%v\" Not Equal %v\n", needCheck, correct)
	}
}

func TestKgToLb(t *testing.T) {
	n := randbtw.Generate(1, 9999)
	t.Logf("Generate-Random-Integer: %v\n", n)

	correct := Pound(float64(n) * 2.2046226218) // correct formula
	needCheck := KgToLb(Kilogram(n))

	if needCheck != correct {
		t.Errorf("\"%v\" Not Equal %v\n", needCheck, correct)
	}
}
