package wtconv

import (
	"fmt"
)

type Lb float64
type Kg float64

func (lb Lb) String() string {
	return fmt.Sprintf("%g lb", lb)
}

func (kg Kg) String() string {
	return fmt.Sprintf("%g kg", kg)
}
