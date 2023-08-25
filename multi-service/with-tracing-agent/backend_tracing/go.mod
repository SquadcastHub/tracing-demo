module main

go 1.20

require (
	github.com/go-chi/chi/v5 v5.0.10
	github.com/gomodule/redigo v1.8.9
	github.com/riandyrn/otelchi v0.5.1
	github.com/signalfx/splunk-otel-go/instrumentation/github.com/gomodule/redigo/splunkredigo v1.7.0
	go.opentelemetry.io/otel v1.16.0
	go.opentelemetry.io/otel/exporters/jaeger v1.16.0
	go.opentelemetry.io/otel/sdk v1.16.0
	go.opentelemetry.io/otel/trace v1.16.0
)

require (
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/signalfx/splunk-otel-go/instrumentation/internal v1.7.0 // indirect
	go.opentelemetry.io/contrib v1.0.0 // indirect
	go.opentelemetry.io/otel/metric v1.16.0 // indirect
	golang.org/x/sys v0.8.0 // indirect
)
