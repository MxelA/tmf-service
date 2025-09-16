package core

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/joho/godotenv"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"log"
	"os"
)

type Tracer struct {
	TraceProvider *sdktrace.TracerProvider
}

func NewJaegerTracer(l *Logger) *Tracer {
	_ = godotenv.Load()

	url, ok := os.LookupEnv("JAEGER_COLLECTOR_HOST")
	if !ok {
		log.Fatalf(".env property JAEGER_COLLECTOR_HOST not found")
	}

	exporter, err := otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure(), otlptracehttp.WithEndpoint(url))
	if err != nil {
		log.Fatal(err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("tmf-service"),
		)),
	)

	otel.SetTracerProvider(tp)

	l.GetCore().Info("Initialize Jaeger Tracer tmf-service")
	//return tp.Shutdown
	return &Tracer{
		TraceProvider: tp,
	}
}

func (jt *Tracer) GetCore() *Tracer {
	return jt
}

func (jt *Tracer) InjectTraceInEventMessage(ctx context.Context, msg *message.Message) {
	otel.SetTextMapPropagator(propagation.TraceContext{})

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	if msg.Metadata == nil {
		msg.Metadata = make(message.Metadata)
	}

	for k, v := range carrier {
		msg.Metadata[k] = v
	}
}

func (jt *Tracer) ExtractTraceFromEventMessage(msg *message.Message) context.Context {
	carrier := propagation.MapCarrier(msg.Metadata)
	return otel.GetTextMapPropagator().Extract(context.Background(), carrier)
}

func (jt *Tracer) Trace(ctx context.Context, spanName string, tracerName string, next func(ctx context.Context) error) (context.Context, error) {

	tracer := otel.Tracer(tracerName)
	ctx, span := tracer.Start(ctx, spanName)
	defer span.End()

	err := next(ctx)

	return ctx, err
}
