package infake

import (
	"fmt"
	"io"
	"os"

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

func (w IoWriterConsumer) Consume(pts <-chan Point) error {
	for p := range pts {
		fmt.Fprintf(w.Writer, "%s\n", p)
	}

	return nil
}

type InfluxDBConsumer struct {
	Client            client.Client
	BatchPointsConfig client.BatchPointsConfig
	BatchSize         uint

	batchPoints client.BatchPoints
}

func (w *InfluxDBConsumer) flush() error {
	if w.batchPoints == nil {
		err := w.makeBatchPoints()

		if err != nil {
			return err
		}
	}

	err := w.Client.Write(w.batchPoints)

	if err != nil {
		return err
	}

	return w.makeBatchPoints()
}

func (w *InfluxDBConsumer) makeBatchPoints() error {
	batchPoints, err := client.NewBatchPoints(w.BatchPointsConfig)

	if err != nil {
		return err
	}

	w.batchPoints = batchPoints

	return nil
}

func (w InfluxDBConsumer) Consume(pts <-chan Point) error {
	if w.batchPoints == nil {
		err := w.makeBatchPoints()

		if err != nil {
			return err
		}
	}

	var consumed uint

	for p := range pts {
		if consumed >= w.BatchSize && w.BatchSize > 0 {
			err := w.flush()

			if err != nil {
				return err
			}

			consumed = 0
		}

		w.batchPoints.AddPoint(p.Point)

		consumed += 1
	}

	return w.flush()
}
