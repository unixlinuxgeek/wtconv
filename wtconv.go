package wtconv

import "fmt"

type Pound float64
type Kilogram float64

func (lb Pound) String() string {
	return fmt.Sprintf("%flb.", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%fkg.", kg)
}
