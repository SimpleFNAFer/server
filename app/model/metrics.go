package model

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	LastSecondRequests prometheus.GaugeFunc
}
