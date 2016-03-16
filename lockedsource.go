package infake

import (
	"math/rand"
	"sync"
)

// LockedSource is a wrapper around a rand.Source which synchronizes access
type LockedSource struct {
	src   rand.Source
	mutex sync.Mutex
}

func (s *LockedSource) Int63() int64 {
	s.mutex.Lock()
	v := s.src.Int63()
	s.mutex.Unlock()
	return v
}

func (s *LockedSource) Seed(seed int64) {
	s.mutex.Lock()
	s.src.Seed(seed)
	s.mutex.Unlock()
}
