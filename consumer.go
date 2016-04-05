package infake

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/influxdata/influxdb/client/v2"
)

type Consumer interface {
	Consume(<-chan Point) error
}

func NewConsumer(cfg OutputConfig) (Consumer, error) {
	switch cfg.Type {
	case "", "stdout":
		return IoWriterConsumer{os.Stdout}, nil
	case "http", "udp":
		var c InfluxDBConsumer
		var err error

		c = InfluxDBConsumer{
			BatchPointsConfig: cfg.BatchPoints,
			BatchSize:         cfg.BatchSize,
			MaxConcurrency:    cfg.MaxConcurrency,
		}

		if cfg.Type == "http" {
			c.Client, err = client.NewHTTPClient(cfg.HTTP)
		} else if cfg.Type == "udp" {
			c.Client, err = client.NewUDPClient(cfg.UDP)
		}

		if err != nil {
			return nil, err
		}

		return c, nil
	}

	return nil, fmt.Errorf("unknown output type: %q", cfg.Type)
}

type IoWriterConsumer struct {
	io.Writer
}

func (c IoWriterConsumer) Consume(pts <-chan Point) error {
	for p := range pts {
		fmt.Fprintf(c.Writer, "%s\n", p)
	}

	return nil
}

type InfluxDBConsumer struct {
	Client            client.Client
	BatchPointsConfig client.BatchPointsConfig
	BatchSize         uint
	MaxConcurrency    uint
}

func (c InfluxDBConsumer) Consume(pts <-chan Point) error {
	bps := make(chan client.BatchPoints)
	errc := make(chan error, 1)

	defer close(errc)

	go func() {
		defer close(bps)

		var bp client.BatchPoints
		var consumed uint
		var err error

		batchSize := c.BatchSize

		if batchSize < 1 {
			batchSize = 1
		}

		for p := range pts {
			if bp == nil {
				bp, err = client.NewBatchPoints(c.BatchPointsConfig)

				if err != nil {
					errc <- err
					return
				}
			}

			if consumed >= batchSize {
				bps <- bp

				bp = nil

				consumed = 0
			} else {
				bp.AddPoint(p.Point)

				consumed += 1
			}
		}

		errc <- nil
	}()

	var maxConcurrency int

	if c.MaxConcurrency < 1 {
		maxConcurrency = 1
	} else {
		maxConcurrency = int(c.MaxConcurrency)
	}

	var wg sync.WaitGroup
	wg.Add(maxConcurrency)

	for i := 0; i < maxConcurrency; i++ {
		go func() {
			c.consumeBatchPoints(bps)
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
	}()

	return <-errc
}

func (c InfluxDBConsumer) consumeBatchPoints(bps <-chan client.BatchPoints) {
	for bp := range bps {
		err := c.Client.Write(bp)

		if err != nil {
			log.Println(err)
		}
	}
}
