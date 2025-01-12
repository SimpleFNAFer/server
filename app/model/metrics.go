package model

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
)

type Metrics struct {
	sync.Mutex

	LastSecondRequests *prometheus.GaugeVec
}
