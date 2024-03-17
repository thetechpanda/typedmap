# TypedMap

[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/typedmap)](https://goreportcard.com/report/github.com/thetechpanda/typedmap)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/typedmap.svg)](https://pkg.go.dev/github.com/thetechpanda/typedmap)
[![Release](https://img.shields.io/github/release/thetechpanda/typedmap.svg?style=flat-square)](https://github.com/thetechpanda/typedmap/tags/latest)

`TypedMap` implements a simple thread-safe map that behaves similarly to `sync.Map` adding type safety and making it simple to know how many unique keys are in the map, `TypeMap` implements the same interface as `sync.Map` and provides `Map[K, V]` interface to ease code refactoring from `sync.Map` to `TypedMap`.

While `TypedMap` provides some similar functionality to `sync.Map`, it is not a drop-in replacement as at is core it uses RWMutex, while `sync.Map` has a more specialised use case. See [sync.Map](https://pkg.go.dev/sync#Map) for its specific use cases.

Consider the following when using `TypedMap` with pointer values:

* **Concurrent Modification:** If multiple goroutines modify the data pointed to by the same pointer without proper synchronization, it can lead to race conditions and unpredictable behavior.
* **Data Race:** Even if the `TypedMap` itself is thread-safe, the data pointed to by the values is not automatically protected. Accessing or modifying the data through pointers in concurrent goroutines can cause data races.

## Documentation

You can find the generated go doc [here](godoc.txt).

## Key Features

* **Type Safety:** Uses generics to provide a type-safe interface for keys and values, eliminating the need for specialised structs and interfaces.
* **Thread Safety:** Ensures safe concurrent access to the map through the use of a sync.RWMutex.
* **Atomic Updates:** Includes functions that allows for atomic modifications to values in the map.
* **Iteration:** Supports iterating over the map with the Range function, and provides methods to obtain slices of keys (Keys), values (Values), or both (Entries).
* **Map Size:** Offers a Len function to easily retrieve the number of items in the map.

## Motivation

Go's `sync.Map` has a general-purpose design, optimized for use cases where keys are dynamically added and removed by multiple goroutines. 

However, this flexibility comes at the cost of type safety and can lead to more verbose and complex code. 

`TypedMap`, on the other hand, leverages generics to offer a more streamlined and type-safe interface, making it easier to work with while still providing the necessary thread safety for concurrent operations.

## Usage

```go
package main

import (
	"fmt"
	"github.com/thetechpanda/typedmap"
)

func main() {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	// v is int
}
```

## Benchmarks
The benchmarks aim to compare the performance of `TypedMap` with that of `sync.Map`. Since `sync.Map` is not typed, the benchmarks focus on comparing the performance of operations that are common to both maps.

```
goos: darwin
goarch: arm64
pkg: github.com/thetechpanda/typedmap
BenchmarkConcurrentSyncMapStore-12                305601              4163 ns/op             347 B/op          7 allocs/op
BenchmarkConcurrentTypedMapSet-12                 603786              2143 ns/op             128 B/op          2 allocs/op
BenchmarkSyncMapDelete-12                        2164330               600.7 ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapDelete-12                       2236756               575.0 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapRange-12                         3766820               386.8 ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapRange-12                       16939611                75.03 ns/op            0 B/op          0 allocs/op
BenchmarkSyncMapLoad-12                          2006686               594.0 ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapGet-12                          2791003               465.5 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapSimulateEntries-12               2137063               505.0 ns/op            96 B/op          0 allocs/op
BenchmarkTypedMapEntries-12                     19449562                63.81 ns/op           16 B/op          0 allocs/op
BenchmarkSyncMapSimulateKeys-12                  3542002               427.2 ns/op            95 B/op          0 allocs/op
BenchmarkTypedMapKeys-12                        29192817                56.57 ns/op            8 B/op          0 allocs/op
BenchmarkSyncMapSimulateValues-12                2027767               509.1 ns/op           101 B/op          0 allocs/op
BenchmarkTypedMapValues-12                      21905688                53.77 ns/op            8 B/op          0 allocs/op
BenchmarkSyncMapSimulateUpdate-12                1000000              1254 ns/op              31 B/op          1 allocs/op
BenchmarkTypedMapUpdate-12                       1975900               624.5 ns/op            16 B/op          1 allocs/op
BenchmarkSyncMapSimulateUpdateRange-12           1373944               941.0 ns/op            31 B/op          1 allocs/op
BenchmarkTypedMapUpdateRange-12                 10460660               114.7 ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapConcurrentOperations-12                9         116667810 ns/op         3200017 B/op     175688 allocs/op
BenchmarkTypedMapConcurrentOperations-12               4         360757323 ns/op           25170 B/op        231 allocs/op
BenchmarkSyncMapConcurrentStore-12                    10         212877171 ns/op         2818623 B/op     174783 allocs/op
BenchmarkTypedMapConcurrentStore-12                   14          96019110 ns/op           21964 B/op        202 allocs/op
BenchmarkSyncMapConcurrentSwap-12                      9         216667921 ns/op         2821499 B/op     174819 allocs/op
BenchmarkTypedMapConcurrentSwap-12                    13          99602955 ns/op           22547 B/op        204 allocs/op
BenchmarkSyncMapConcurrentLoadOrStore-12               5         220694367 ns/op         4113673 B/op     162350 allocs/op
BenchmarkTypedMapConcurrentLoadOrStore-12              8         134787792 ns/op           19821 B/op        214 allocs/op
BenchmarkTypedMapConcurrentUpdate-12                  14          99482089 ns/op           21439 B/op        197 allocs/op
PASS
ok      github.com/thetechpanda/typedmap        166.434s
```

### Concurrent Benchmarks
Concurrent benchmark use all the same function body, so that each bench has the behaviour except for TypedMap and sync.Map operations.

Check `benchmarkConcurrentInt` to check how the bench behaves

## Installation

```bash
go get github.com/thetechpanda/typedmap
```

## Contributing

Contributions are welcome and very much appreciated! 

Feel free to open an issue or submit a pull request.

## License

`TypedMap` is released under the MIT License. See the [LICENSE](LICENSE) file for details.
