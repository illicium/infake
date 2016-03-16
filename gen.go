package infake

import (
	"log"
	"sync"

	"math/rand"
)

type Gen struct {
	rndSrc rand.Source
	Config
}

func NewGen(cfg Config) *Gen {
	rndSrc := &LockedSource{src: rand.NewSource(cfg.Seed)}

	return &Gen{rndSrc, cfg}
}

func (gen *Gen) Generate() (<-chan Point, error) {
	var wg sync.WaitGroup
	c := make(chan Point)

	wg.Add(len(gen.Series))

	for _, series := range gen.Series {
		go func(series Series) {
			defer wg.Done()

			sc, err := series.Generate(gen.rndSrc)

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
