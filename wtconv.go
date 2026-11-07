package wtconv

import "fmt"

type Pound float64
type Kilogram float64

const lb = 2.2046226218 // pound

func (lb Pound) String() string {
	return fmt.Sprintf("%flb.", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%fkg.", kg)
}
