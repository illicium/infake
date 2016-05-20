package infake

import (
	"fmt"
	"math/rand"
)

const DEFAULT_INT_MIN int = -1000
const DEFAULT_INT_MAX int = 1000
const DEFAULT_FLOAT_MIN float64 = -1.0
const DEFAULT_FLOAT_MAX float64 = 1.0

type Value interface {
	Get(*rand.Rand) interface{}
}

type ConstantValue struct {
	Value interface{}
}

func (v *ConstantValue) Get(rnd *rand.Rand) interface{} {
	return v.Value
}

type IntValue struct {
	Min int64
	Max int64
}

func (v *IntValue) Get(rnd *rand.Rand) interface{} {
	if v.Max-v.Min == 0 {
		return v.Min
	}
	return rand.Int63n((v.Max-v.Min)+1) + v.Min
}

type FloatValue struct {
	Min float64
	Max float64
}

func (v *FloatValue) Get(rnd *rand.Rand) interface{} {
	return rand.Float64()*(v.Max-v.Min) + v.Min
}

type ExpFloatValue struct {
	Rate float64
}

func (v *ExpFloatValue) Get(rnd *rand.Rand) interface{} {
	return rand.ExpFloat64() / v.Rate
}

type NormFloatValue struct {
	StdDev float64
	Mean   float64
}

func (v *NormFloatValue) Get(rnd *rand.Rand) interface{} {
	return rand.NormFloat64()*v.StdDev + v.Mean
}

type ValueConfig struct {
	Type  string
	Value interface{}

	// Type: int, float
	Min interface{}
	Max interface{}

	// Type: expfloat
	Rate float64

	// Type: normfloat
	StdDev float64
	Mean   float64
}

func newIntValue(c ValueConfig) (Value, error) {
	var min int
	var max int
	var ok bool

	if c.Min == nil {
		min = DEFAULT_INT_MIN
	} else if min, ok = c.Min.(int); !ok {
		return nil, fmt.Errorf("bad int min value: %v", c.Min)
	}

	if c.Max == nil {
		max = DEFAULT_INT_MAX
	} else if max, ok = c.Max.(int); !ok {
		return nil, fmt.Errorf("bad int max value: %v", c.Max)
	}

	return &IntValue{int64(min), int64(max)}, nil
}

func newFloatValue(c ValueConfig) (Value, error) {
	var min float64
	var max float64
	var ok bool

	if c.Min == nil {
		min = DEFAULT_FLOAT_MIN
	} else if min, ok = c.Min.(float64); !ok {
		var minInt int
		if minInt, ok = c.Min.(int); ok {
			min = float64(minInt)
		} else {
			return nil, fmt.Errorf("bad float min value: %v", c.Min)
		}
	}

	if c.Max == nil {
		max = DEFAULT_FLOAT_MAX
	} else if max, ok = c.Max.(float64); !ok {
		var maxInt int
		if maxInt, ok = c.Max.(int); ok {
			max = float64(maxInt)
		} else {
			return nil, fmt.Errorf("bad float max value: %v", c.Max)
		}
	}

	return &FloatValue{float64(min), float64(max)}, nil
}

func newExpFloatValue(c ValueConfig) (Value, error) {
	rate := c.Rate

	if rate == 0 {
		rate = 1.0
	}

	return &ExpFloatValue{rate}, nil
}

func newNormFloatValue(c ValueConfig) (Value, error) {
	stdDev := c.StdDev

	if stdDev == 0 {
		stdDev = 1.0
	}

	return &NormFloatValue{stdDev, c.Mean}, nil
}

func NewValue(c ValueConfig) (Value, error) {
	if c.Value != nil {
		return &ConstantValue{c.Value}, nil
	}

	switch c.Type {
	case "int": // int64
		return newIntValue(c)
	case "float": // float64
		return newFloatValue(c)
	case "expfloat": // float64
		return newExpFloatValue(c)
	case "normfloat": // float64
		return newNormFloatValue(c)
	}
	// TODO: string, bool

	return nil, fmt.Errorf("unknown value type: %s", c.Type)
}
