# Benchmarking and verifying rate limit behaviour

Currently this does the following:

- Start a gateway with a rate limiter configuration flag via env (`task leaky`).
- Issue hey at a request rate of 50 requests/s for 3 seconds.
- Log the incoming request rate and HTTP status code responses with rakyll/hey.
- Log the back-end request rate and HTTP responses (200 OK, can assert JSON).
- Log the gateway `--memoryprofile` during the benchmark.
- Log the gateway `--cpuprofile` during the benchmark.

From these `.json` and `.csv` files, a report should be generated.

TODO:

- [ ] declare an api with a `rate=40` and `per=1` value,
- [ ] replicate test for other rate limit algos
- [ ] sanitize and analyze data sources, compare
