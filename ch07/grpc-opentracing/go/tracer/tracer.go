package tracer

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	metricsFactory := prometheus.New()
	reporterConfig := config.ReporterConfig{
		QueueSize:           1000,
		BufferFlushInterval: 2000,
		LogSpans:            true,
		LocalAgentHostPort:  "localhost/5775",
	}
	samplerConfig := config.SamplerConfig{
		Type:  "const",
		Param: 1.0,
	}
	return config.Configuration{
		ServiceName: servicename,
		Disabled:    false,
		Reporter:    &reporterConfig,
		Sampler:     &samplerConfig,
	}.NewTracer(
		config.Metrics(metricsFactory),
	)
}
