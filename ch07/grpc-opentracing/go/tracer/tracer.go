package tracer

import (
	"github.com/uber/jaeger-client-go"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics/prometheus"
)

func NewTracer(servicename string) (opentracing.Tracer, io.Closer, error) {
	// load config from environment variables
	cfg := config.Configuration{
		ServiceName: servicename,
		Sampler: &config.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "127.0.0.1:6831",
		},
	}

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
	jLogger := log.StdLogger
	metricsFactory := prometheus.New()
	return cfg.NewTracer(
		config.Logger(jLogger),
		config.Metrics(metricsFactory),
	)
}
