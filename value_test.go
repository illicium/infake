package infake

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

type SourceMock struct {
	Value     int64
	SeedValue int64 // unused
}

func (s *SourceMock) Int63() int64 {
	return s.Value
}

func (s *SourceMock) Seed(seed int64) {
	s.SeedValue = seed
}

func getMockSource() rand.Source {
	return &SourceMock{}
}

func getMockRand() *rand.Rand {
	return rand.New(getMockSource())
}

func TestConstantValue(t *testing.T) {
	rnd := getMockRand()

	v := ConstantValue{123}

	assert.Equal(t, 123, v.Get(rnd))
}
