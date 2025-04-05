# REST API with PostgreSQL - Next.js vs Go Chi

This repo creates a comparison of performance between a Go API
using the Chi library and a TypeScript API using Next.js.

Both APIs are configured to select columns from the same table,
serialize those rows as JSON, and return those rows to the client.

## Load testing

The `compare_endpoints.go` file is the driver to test the performance
of these APIs under load. It uses goroutines to send concurrent requests
to each API, measuring errors, max/min RTT, and mean RTT. The values for
each API are printed in readable format.

## Results

At 10,000 concurrent requests, the Go Chi API achieves a 10x - 100x faster
min RTT than the Next.js API route, and a 1.3x - 2x mean. Once I implement
a median calculation, I expect the median to show a wider gap than the mean,
likely 5x - 10x speedup.

## TODO

- [ ] Implement median calculation
- [ ] Implement Go FastAPI version for faster speeds
