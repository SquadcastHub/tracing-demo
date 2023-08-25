# tracing-demo 

A demo project to showcase jaeger-tracing.

## What is distributed tracing?

Distributed tracing is a method of monitoring and observing service requests in applications built on a microservices architecture.

It tracks and traces requests as they flow through distributed systems or microservice environments, allowing developers and devops teams to gain end-to-end observability into the journey of each request.



:smiley: **Benifits of Tracing** 

* Provides end-to-end visibility into the journey of each request, allowing for better troubleshooting and debugging of performance issues
* Helps pinpoint where failures occur and what causes suboptimal performance
* Enables tracking of application requests as they flow across service boundaries, providing a complete picture of the request trace

:sweat_smile: **Challenges**

* Generating and collecting trace data in high-throughput systems can be challenging, as millions of spans may be generated per minute.
* Managing telemetry clusters.
* Code changes required to enable traces.


## The three pillars of observability

**Tracing**

* Tracing provides visibility into how a request is processed across multiple services in a microservices environment.
* Tracing helps identify the root cause of an issue by following the execution path of a request and recording the time taken for each service to respond.
* Traces provide a detailed view of the application's flow and data progression, allowing developers to pinpoint specific applications or services causing problems.
* Tracing is especially useful for end-to-end monitoring and debugging of defects.

**Logging**
* Logging is the process of tracking detailed information about events in an application, particularly errors, warnings, or other exceptional situations.
* Logs provide a record of what happened in the system, allowing developers to investigate issues and troubleshoot problems.
* Logs are essential for application monitoring and should always be enabled.
* Logging is useful for capturing specific events and can be used in conjunction with tracing to understand the root cause of an issue.

**Metrics**
* Metrics are used to track the occurrence of events, count items, measure the time taken to perform actions, or report the current value of resources like CPU or memory.
* Metrics provide a high-level overview of the system's performance and behavior at defined time intervals.
* Metrics are useful for monitoring and alerting purposes, as they can help set thresholds and trigger alerts when certain conditions are met.
* Metrics do not provide a detailed view of the application's execution path or individual requests like tracing does.

In summary, tracing, logging, and metrics are all important for observability in an application, but they serve different purposes. Tracing helps identify the root cause of issues by following the execution path of a request, logging captures detailed information about events, and metrics provide a high-level overview of the system's performance.


## Tools for Capturing Traces.

- ZIPKIN
- Jaeger 
- Opentelmetry 
- GCP traces
- AWS X-RAY

## What is Opentelmetry and Jaeger? :thinking:

**Opentelmetry**

OpenTelemetry is an open-source observability framework designed to provide a unified set of APIs, libraries, agents, and instrumentation for capturing telemetry data from applications.  

OpenTelemetry aims to provide a comprehensive approach to observability by including tracing, metrics, and logging as part of its framework. 

**Jaeger** is a project within the OpenTelemetry ecosystem and can be used as the default backend when using OpenTelemetry to gather telemetry data from applications

It is a specialized open-source distributed tracing system, originally developed by Uber Technologies. It helps gather timing data for requests as they flow through a distributed system, allowing developers to analyze and debug performance issues. Jaeger provides a web-based user interface for visualizing traces and supports high scalability for large-scale tracing.

There are some other alternatives for jaeger as well like **new relic**, **data dog**, **lightstep** etc..

## Some Terminlogies before we start demo

* **Tracing** Traces give us the big picture of what happens when a request is made to an application. Whether your application is a monolith with a single database or a sophisticated mesh of services, traces are essential to understanding the full “path” a request takes in your application.
* **Spans** A span represents a logical unit of work that has an operation name, the start time of the operation, and the duration. Spans may be nested and ordered to model causal relationships.
* A **Tracer Provider** (sometimes called TracerProvider) is a factory for Tracers. In most applications, a Tracer Provider is initialized once and its lifecycle matches the application’s lifecycle.
* **Tracer** A Tracer creates spans containing more information about what is happening for a given operation, such as a request in a service. Tracers are created from Tracer Providers.
* **Trace Exporters** Trace Exporters send traces to a consumer. This consumer can be standard output for debugging and development-time, the OpenTelemetry Collector,Jaeger exporter or any open source or vendor backend of your choice.


* **Context Propagation**

Context Propagation is the core concept that enables Distributed Tracing. With Context Propagation, Spans can be correlated with each other and assembled into a trace, regardless of where Spans are generated. We define Context Propagation by two sub-concepts: Context and Propagation.

A **Context** is an object that contains the information for the sending and receiving service to correlate one span with another and associate it with the trace overall. For example, if Service A calls Service B, then a span from Service A whose ID is in context will be used as the parent span for the next span created in Service B.

**Propagation** is the mechanism that moves context between services and processes. It serializes or deserializes the context object and provides the relevant Trace information to be propagated from one service to another.

---


* **Instrumentation Library** Denotes the library that provides the instrumentation for a given Instrumented Library.


