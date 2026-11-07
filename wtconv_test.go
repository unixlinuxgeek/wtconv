package wtconv

import (
	"testing"

	"github.com/unixlinuxgeek/randbtw"
)

func TestLbToKg(t *testing.T) {
	n := randbtw.Generate(1, 9999)
	t.Logf("Generated-Random-Integer: %v\n", n)

	correct := Kilogram(float64(n) / Lb) // correct formula
	needCheck := LbToKg(Pound(n))

	if testing.Verbose() {
		t.Logf("correct: %v\n", correct)
		t.Logf("LbToKg(%v)\n", needCheck)
	}

	if needCheck != correct {
		t.Errorf("\"%v\" Not Equal %v\n", needCheck, correct)
	}
}

func TestKgToLb(t *testing.T) {
	n := randbtw.Generate(1, 9999)
	t.Logf("Generated-Random-Integer: %v\n", n)

	correct := Pound(float64(n) * Lb) // correct formula
	needCheck := KgToLb(Kilogram(n))

	if testing.Verbose() {
		t.Logf("correct: %v\n", correct)
		t.Logf("KgToLb(%v)\n", needCheck)
	}

	if needCheck != correct {
		t.Errorf("\"%v\" Not Equal %v\n", needCheck, correct)
	}
}
