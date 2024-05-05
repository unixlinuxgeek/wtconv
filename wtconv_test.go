package wtconv

import (
	"github.com/unixlinuxgeek/floatgen"
	"testing"
)

// KgToLb Умножаем значение массы на 2.205
func TestKgToLb(t *testing.T) {
	kg := Kg(floatgen.GenRan(1, 9))
	lb := Lb(kg * 2.205)

	if lb == KgToLb(kg) {
		t.Logf("TestKgToLb Passed!!!: %v equal %v\n", kg, lb)
	} else {
		t.Fatalf("TestKgToLb Failed: %v not equal %v\n", kg, lb)
	}
}

// LbToKg Разделим значение массы на 2.205
func TestLbToKg(t *testing.T) {
	lb := Lb(floatgen.GenRan(1, 9))
	kg := Kg(lb / 2.205)

	if kg == LbToKg(lb) {
		t.Logf("TestLbToKg Passed!!!: %v equal %v\n", lb, kg)
	} else {
		t.Fatalf("TestLbToKg Failed: %v not equal %v\n", lb, kg)
	}
}
