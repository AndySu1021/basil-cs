package tool

import (
	"cs-api/pkg/types"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

type RequestInstrument struct {
	RequestCount   *prometheus.CounterVec
	RequestLatency *prometheus.HistogramVec
}

func NewRequestInstrument() *RequestInstrument {
	labels := []string{"op"}

	counter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: types.RequestCount,
		Help: types.RequestCountHelp,
	}, labels)

	latency := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: types.RequestLatency,
		Help: types.RequestLatencyHelp,
	}, labels)

	prometheus.MustRegister(counter, latency)

	return &RequestInstrument{
		RequestCount:   counter,
		RequestLatency: latency,
	}
}

func (i *RequestInstrument) Op(op string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func(begin time.Time) {
			i.RequestCount.WithLabelValues(op).Add(1)
			i.RequestLatency.WithLabelValues(op).Observe(time.Since(begin).Seconds())
		}(time.Now())
	}
}
