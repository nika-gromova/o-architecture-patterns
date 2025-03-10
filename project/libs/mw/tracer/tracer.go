package tracer

import (
	"context"

	"github.com/nika-gromova/o-architecture-patterns/project/libs/tracer"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"google.golang.org/grpc"
)

func InterceptorGRPC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		span, ctx = opentracing.StartSpanFromContext(ctx, info.FullMethod)
		defer span.Finish()
	}

	if spanContext, ok := span.Context().(jaeger.SpanContext); ok {
		_ = grpc.SendHeader(ctx, map[string][]string{
			"x-trace-id": {spanContext.TraceID().String()},
		})
	}
	h, err := handler(ctx, req)

	if err != nil {
		err = tracer.MarkSpanWithError(ctx, err)
	}

	return h, err
}
