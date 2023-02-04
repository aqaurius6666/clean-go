package otel

import (
	"context"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/trace"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	metricglobal "go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OtelOptions struct {
	CollectorAddr  string
	ID             int64
	ServiceName    string
	MetricPeriodic time.Duration
}

func InitOtel(ctx context.Context, opts OtelOptions) (func(ctx context.Context) error, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, opts.CollectorAddr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	metricExporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithGRPCConn(conn),
	)
	if err != nil {
		return nil, err
	}
	metricProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.ServiceName),
			attribute.Int64("ID", opts.ID),
		)),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(opts.MetricPeriodic),
			),
		),
	)
	metricglobal.SetMeterProvider(metricProvider)
	err = runtime.Start(runtime.WithMinimumReadMemStatsInterval(opts.MetricPeriodic))
	if err != nil {
		return nil, err
	}
	traceExporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithGRPCConn(conn),
	)
	if err != nil {
		return nil, err
	}
	c, _ := metricglobal.Meter("asd").SyncInt64().Counter("asd")
	c.Add(ctx, 1)
	bsp := tracesdk.NewBatchSpanProcessor(traceExporter)
	traceProvider := tracesdk.NewTracerProvider(
		tracesdk.WithSampler(tracesdk.AlwaysSample()),
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(opts.ServiceName),
			attribute.Int64("ID", opts.ID),
		)),
		tracesdk.WithSpanProcessor(bsp),
	)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	otel.SetTracerProvider(traceProvider)

	return func(ctx context.Context) error {
		err = metricExporter.Shutdown(ctx)
		if err != nil {
			return err
		}

		err = traceExporter.Shutdown(ctx)
		if err != nil {
			return err
		}
		return nil
	}, err
}

func TracerProvider(name string, opts ...trace.TracerOption) trace.TracerProvider {
	return otel.GetTracerProvider()
}

var NoOpTracer = trace.NewNoopTracerProvider()
