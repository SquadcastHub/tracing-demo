package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/riandyrn/otelchi"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	version              = "unknown"
	envKeyBackServiceURL = "BACKEND_SVC_URL"
	serviceName          = "front-svc-with-agent"
)

func NewTracer(svcName string) (trace.Tracer, error) {
	// create jaeger exporter
	exporter, err := jaeger.New(jaeger.WithAgentEndpoint(jaeger.WithAgentHost("localhost"), jaeger.WithAgentPort("6831")))
	if err != nil {
		return nil, fmt.Errorf("unable to initialize exporter due: %w", err)
	}
	// initialize tracer provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(svcName),
		)),
	)
	// set tracer provider and propagator properly, this is to ensure all
	// instrumentation library could run well
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// returns tracer
	return otel.Tracer(svcName), nil
}

func main() {
	addr := os.Getenv("FRONTEND_APP_ADDR")
	tracer, err := NewTracer(serviceName)
	if err != nil {
		log.Fatalf("unable to initialize tracer due: %v", err)
	}

	r := chi.NewRouter()
	r.Use(otelchi.Middleware(serviceName, otelchi.WithChiRoutes(r)))

	r.Get("/counter/{keyspace}", func(w http.ResponseWriter, r *http.Request) {
		keyspace := chi.URLParam(r, "keyspace")
		count, err := getCount(keyspace, r.Context(), tracer)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.Write([]byte(fmt.Sprintf("Counter for %s: %s", keyspace, count)))
	})
	// execute server
	log.Printf("front service is listening on %v", addr)
	err = http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatalf("unable to execute server due: %v", err)
	}
}

func getCount(keyspace string, ctx context.Context, tracer trace.Tracer) (string, error) {
	new_ctx, span := tracer.Start(ctx, "getCount")
	defer span.End()

	requestURL := fmt.Sprintf("http://%s/%s", os.Getenv(envKeyBackServiceURL), keyspace)
	resp, err := otelhttp.Get(new_ctx, requestURL)
	defer resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("unable to execute http request due: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("unable to read response data due: %w", err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return "", err
	}
	sb := string(body)

	return sb, nil
}
