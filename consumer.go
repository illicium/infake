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

		c = InfluxDBConsumer{}

		if cfg.Type == "http" {
			c.Client, err = client.NewHTTPClient(cfg.HTTP)
		} else if cfg.Type == "udp" {
			c.Client, err = client.NewUDPClient(cfg.UDP)
		}

		if err != nil {
			return nil, err
		}

		c.BatchPoints, err = client.NewBatchPoints(cfg.BatchPoints)

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
	Client      client.Client
	BatchPoints client.BatchPoints
}

func (w InfluxDBConsumer) Consume(pts <-chan Point) error {
	for p := range pts {
		w.BatchPoints.AddPoint(p.Point)
	}

	return w.Client.Write(w.BatchPoints)
}
