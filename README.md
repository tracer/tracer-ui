# tracer-ui

tracer-ui provides a web frontend to the query API of the Tracer
distributed tracing system.

## Quick start

Run `tracer-ui -t <path to zipkin-ui folder>` and open your browser at
http://localhost:9997.

See `tracer-ui -h` for a list of all available flags.

## Requirements

The current version of tracer-ui relies on zipkin-ui for rendering
traces. As such, it needs the zipkinhttp query transport of Tracer to
be enabled and accessible.
