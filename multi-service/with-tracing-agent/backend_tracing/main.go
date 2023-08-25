package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/gomodule/redigo/redis"
	"github.com/riandyrn/otelchi"
	splunkredis "github.com/signalfx/splunk-otel-go/instrumentation/github.com/gomodule/redigo/splunkredigo/redis"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var (
	// injected during build
	version     = "unknown"
	serviceName = "back-svc-with-agent"
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

// initCachePool initializes redis for cache
func initCachePool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			return splunkredis.Dial("tcp", addr)
		},
	}
}

func main() {
	// init redis
	cachePool := initCachePool(os.Getenv("DEMO_REDIS_ADDR"))

	// check if redis is alive or not
	conn := cachePool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		panic(fmt.Sprintf("error initializing cache pool: %v", err))
	}

	// Initialise tracer
	tracer, err := NewTracer(serviceName)
	if err != nil {
		log.Fatalf("unable to initialize tracer provider due: %v", err)
	}

	// initialise handlers
	r := chi.NewRouter()

	r.Use(otelchi.Middleware(serviceName, otelchi.WithChiRoutes(r)))

	r.Get("/{keyspace}", func(w http.ResponseWriter, r *http.Request) {
		keyspace := chi.URLParam(r, "keyspace")
		err = incrementKey(conn, keyspace, r.Context(), tracer)
		if err != nil {
			w.Write([]byte(fmt.Sprintf("oops something went wrong: %v", err)))
			return
		}
		val, err := redis.Int(conn.Do("GET", keyspace, r.Context()))
		if err != nil {
			w.Write([]byte(fmt.Sprintf("oops something went wrong: %v", err)))
			return
		}
		w.Write([]byte(fmt.Sprintf("%d", val)))
	})
	addr := os.Getenv("DEMO_APP_ADDR")
	if addr == "" {
		addr = ":8000"
	}

	log.Printf("Booting app on %s", addr)
	http.ListenAndServe(addr, r)
}

func incrementKey(c redis.Conn, keyspace string, ctx context.Context, tracer trace.Tracer) error {
	new_ctx, span := tracer.Start(ctx, "incrementKey")
	defer span.End()

	c.Do("INCR", keyspace, new_ctx)
	return nil
}
