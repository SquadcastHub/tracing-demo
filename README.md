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
