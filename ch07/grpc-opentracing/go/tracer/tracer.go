package tracer

import (
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	metricsFactory := prometheus.New()
	// load config from environment variables
	cfg, _ := config.FromEnv()
	cfg.ServiceName = servicename
	//reporterConfig := config.ReporterConfig{
	//	QueueSize:           1000,
	//	BufferFlushInterval: 2000,
	//	LogSpans:            true,
	//	LocalAgentHostPort:  "localhost:6831",
	//}
	//samplerConfig := config.SamplerConfig{
	//	Type:  "const",
	//	Param: 1.0,
	//}
	return cfg.NewTracer(
		config.Metrics(metricsFactory),
	)
}
