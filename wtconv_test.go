package wtconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

// KgToLb Умножаем значение массы на 2.205
func TestKgToLb(t *testing.T) {
	kg := floatgen.GenRan(1, 9)
	lb := Lb(kg * 2.205)

	if lb == KgToLb(Kg(kg)) {
		t.Logf("TestKgToLb Passed!!!: %v equal %v\n", kg, lb)
	} else {
		t.Fatalf("TestKgToLb Failed: %v no equal %v\n", kg, lb)
	}
}
