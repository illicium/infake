package infake

import (
	"log"
	"sync"

	"math/rand"
)

type Gen struct {
	rnd *rand.Rand
	Config
}

func NewGen(cfg Config) *Gen {
	rndSrc := rand.NewSource(cfg.Seed)
	rnd := rand.New(rndSrc)

	gen := &Gen{
		rnd,
		cfg,
	}

	return gen
}

func (gen *Gen) Generate() (<-chan Point, error) {
	var wg sync.WaitGroup
	c := make(chan Point)

	wg.Add(len(gen.Series))

	for _, series := range gen.Series {
		go func(series Series) {
			defer wg.Done()

			sc, err := series.Generate(gen.rnd)

			if err != nil {
				log.Print(err)
				return
			}

			for p := range sc {
				c <- p
			}
		}(series)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	return c, nil
}
