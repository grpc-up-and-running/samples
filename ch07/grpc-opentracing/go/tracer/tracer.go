package tracer

import (
	"io"

	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	metricsFactory := prometheus.New()
	return config.Configuration{
		ServiceName: servicename,
	}.NewTracer(
		config.Metrics(metricsFactory),
	)
}
