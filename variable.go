package infake

import (
	"fmt"
	"math"
)

type Variable struct {
	Name   string
	From   float64
	To     float64
	Step   float64
	Value  Value
	Values []Value
}

type Value interface{}

// Expand generates copies of the variable for each possible Value given by Values
func (v *Variable) expandValues() ([]Variable, error) {
	if len(v.Values) == 0 {
		return nil, fmt.Errorf("%q: Values is empty", v.Name)
	}

	vars := make([]Variable, 0, len(v.Values))

	for _, val := range v.Values {
		newVar := *v
		newVar.Values = nil
		newVar.Value = val

		vars = append(vars, newVar)
	}

	return vars, nil
}

func (v *Variable) expandNumeric() ([]Variable, error) {
	step := v.Step

	if step == 0 {
		step = v.To - v.From
	}

	if v.From == v.To && step == 0 {
		return nil, fmt.Errorf("%q: From == To, Step != 0", v.Name)
	}

	if v.From < v.To && step < 0 {
		return nil, fmt.Errorf("%q: infinite sequence: From < To, Step < 0", v.Name)
	}

	if v.To < v.From && step > 0 {
		return nil, fmt.Errorf("%q: infinite sequence: To < From, Step > 0", v.Name)
	}

	var n int

	if v.From < v.To {
		n = int((v.To-v.From)/math.Abs(step) + 1)
	} else {
		n = int((v.From-v.To)/math.Abs(step) + 1)
	}

	vars := make([]Variable, 0, n)

	for i := v.From; (step > 0 && i <= v.To) || (step < 0 && i >= v.To); i += step {
		newVar := *v
		newVar.Value = i

		vars = append(vars, newVar)
	}

	return vars, nil
}

// Expand generates copies of the variable with for each possible Value given by Min/Max or Values
func (v *Variable) Expand() ([]Variable, error) {
	if v.Value != nil {
		return []Variable{*v}, nil
	}

	if v.Values != nil {
		return v.expandValues()
	}

	return v.expandNumeric()
}
