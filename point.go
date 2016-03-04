package infake

import (
	"time"

	"github.com/influxdata/influxdb/client/v2"
)

type Point struct {
	*client.Point
}

// NewPoint returns a new Point
func NewPoint(name string, tags map[string]string, fields map[string]interface{}, t ...time.Time) (Point, error) {
	var T time.Time

	if len(t) > 0 {
		T = t[0]
	}

	p, err := client.NewPoint(name, tags, fields, T)

	if err != nil {
		return Point{}, err
	}

	return Point{p}, nil
}
