package wtconv

import "fmt"

type Pound float64
type Kilogram float64

const Lb = 2.2046226218 // pound

func (lb Pound) String() string {
	return fmt.Sprintf("%f lb.", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%f kg.", kg)
}
