package infake

import "github.com/influxdata/influxdb/client/v2"

type Config struct {
	Seed int64
	// Time   int64
	Output OutputConfig
	Series []Series
}

type OutputConfig struct {
	Type        string
	HTTP        client.HTTPConfig
	UDP         client.UDPConfig
	BatchPoints client.BatchPointsConfig
}
