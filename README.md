Weasel
------

> Eats gophers for breakfast

Comprehensive instrumentation library for Go applications.

Aims to handle metrics, error-tracking (via logging), and profiling in one
simple package.

Currently only exports profiling information, but the rest is on the way

Works like `Prometheus`, relying on a server-poll of an HTTP endpoint to collect
information. However, unlike `Prometheus`, we can optionally register with the
collector, instead of relying on a service-discovery mechanism. This mechanism
is not part of `weasel`, but is the recommended way to use it.

Profiling
------
Weasel is capable of automatically exporting `pprof` data through the same HTTP
channels it securely exports metrics + error data.  

The profiling code has its origins in [netbug](https://github.com/e-dard/netbug)


LICENSE
------
MIT
