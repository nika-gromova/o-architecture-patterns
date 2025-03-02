package tracer

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go/config"
)

func InitGlobal(service string) error {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "localhost:6831",
		},
	}

	if _, err := cfg.InitGlobalTracer(service); err != nil {
		return err
	}
	return nil
}

func MarkSpanWithError(ctx context.Context, err error) error {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return err
	}

	ext.Error.Set(span, true)
	span.LogKV("error", err.Error())

	return err
}
