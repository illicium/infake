package infake

import (
	"log"
	"sync"

	"math/rand"
)

type Gen struct {
	RndSrc rand.Source
	Series []*Series
}

func NewGen(cfg Config) (*Gen, error) {
	rndSrc := &LockedSource{src: rand.NewSource(cfg.Seed)}

	gen := &Gen{
		RndSrc: rndSrc,
	}
	gen.Series = make([]*Series, len(cfg.Series))
	for i, seriesCfg := range cfg.Series {
		s, err := NewSeries(seriesCfg)

		if err != nil {
			return nil, err
		}

		gen.Series[i] = s
	}

	return gen, nil
}

func (gen *Gen) Generate() (<-chan Point, error) {
	var wg sync.WaitGroup
	c := make(chan Point)

	wg.Add(len(gen.Series))

	for _, series := range gen.Series {
		go func(series *Series) {
			defer wg.Done()

			sc, err := series.Generate(gen.RndSrc)

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
