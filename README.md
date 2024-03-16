# TypedMap

[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/typedmap)](https://goreportcard.com/report/github.com/thetechpanda/typedmap)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/typedmap.svg)](https://pkg.go.dev/github.com/thetechpanda/typedmap)
[![Release](https://img.shields.io/github/release/thetechpanda/typedmap.svg?style=flat-square)](https://github.com/thetechpanda/typedmap/tags/latest)

`TypedMap` enhances the standard `sync.Map` in Go by leveraging generics to provide type safety and a cleaner, simpler interface.

While `TypedMap` provides some similar functionality, it is not a drop-in replacement for `sync.Map` and offers a simpler, more specialized interface. Its main objective is to sacrifice some compute time for type safety and a consistent interface.

## Documentation

I like when documentation is available in the repository, [godocdown](https://github.com/robertkrimen/godocdown) does a great job at making beautiful markdown.

You can find the generated markdown [here](godoc.md).

## Key Features

- **Type Safety:** Eliminates the need for type assertions when retrieving values from the map.
- **Simplified Interface:** Offers a more intuitive and reduced interface for common map operations.
- **Performance:** Aims to match the performance of `sync.Map` as closely as possible.

Please read the documentation as `TypedMap` has some of the same behaviour as `map` and `sync.Map`.

- `Keys()` and `Values()` do not guarantee the order of keys and values to be consistent when called individually.
- `Range()` and `Entries()` return keys and values pairs, but the order of the pairs are provided may not consistent between calls.
- `Len()`, `Keys()`, `Values()`, `Entries()` and `Range()` have the same guarantees as `sync.Map.Range`:
  - They do not necessarily correspond to any consistent snapshot of the Map's contents: no key will be visited more than once, but if the value for any key is stored or deleted concurrently.
  - They may reflect any mapping for that key from any point during the call.
- `Set()`, `Delete()`, `Update()`, `UpdateRange()` and `Clear()` are blocking operations and guarantee that no other goroutines can modify the map during these operations.

## Motivation

Go's `sync.Map` has a general-purpose design, optimized for use cases where keys are dynamically added and removed by multiple goroutines. However, this flexibility comes at the cost of type safety and can lead to more verbose and complex code. `TypedMap`, on the other hand, leverages generics to offer a more streamlined and type-safe interface, making it easier to work with while still providing the necessary thread safety for concurrent operations.

`TypedMap` tries improves code clarity and safety but also introduces additional functionality, such as the `Len()` method for tracking the number of items in the map and the `Update()` and `UpdateRange()` method for atomic updates. 
These enhancements make `TypedMap` a compelling option for developers looking for a more convenient and robust solution for managing concurrent data structures in Go.

## Usage

```go
package main

import (
	"fmt"
	"github.com/thetechpanda/typedmap"
)

func main() {
	m := typedmap.New[string, int]()
	m.Set("key1", 10)
	m.Set("key2", 20)

	value, ok := m.Get("key1")
	if ok {
		fmt.Println("Value:", value)
	}

	fmt.Println("Length:", m.Len())

	m.Update("key2", func(oldValue int, exists bool) int {
		if exists {
			return oldValue + 5
		}
		return 5
	})

	newValue, _ := m.Get("key2")
	fmt.Println("Updated Value:", newValue)
}
```

## Benchmarks
The benchmarks aim to compare the performance of `TypedMap` with that of `sync.Map`. Since `sync.Map` is not typed, the benchmarks focus on comparing the performance of operations that are common to both maps.

```
$ go test -bench=. -benchmem -cpu=4
goos: darwin
goarch: arm64
pkg: github.com/thetechpanda/typedmap
BenchmarkConcurrentSyncMapStore-4                1421858               778.5  ns/op           283 B/op          7 allocs/op
BenchmarkConcurrentTypedMapSet-4                 1358841               866.0  ns/op           227 B/op          7 allocs/op
BenchmarkSyncMapDelete-4                        12863085               137.4  ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapDelete-4                        9828331               171.3  ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapRange-4                         31028598                67.5  ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapRange-4                        28134807               115.1  ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapLoad-4                          14091128               152.8  ns/op             0 B/op          0 allocs/op
BenchmarkTypedMapGet-4                          18840867                88.6  ns/op             0 B/op          0 allocs/op
BenchmarkSyncMapSimulateEntries-4               23477412               252.8  ns/op           172 B/op          0 allocs/op
BenchmarkTypedMapEntries-4                      26126026                65.6  ns/op            16 B/op          0 allocs/op
BenchmarkSyncMapSimulateKeys-4                  27779840               129.7  ns/op            90 B/op          0 allocs/op
BenchmarkTypedMapKeys-4                         27185250                66.2  ns/op             8 B/op          0 allocs/op
BenchmarkSyncMapSimulateValues-4                14924552               104.3  ns/op            94 B/op          0 allocs/op
BenchmarkTypedMapValues-4                       23461269                74.9  ns/op             8 B/op          0 allocs/op
BenchmarkSyncMapSimulateUpdate-4                 5814201               245.6  ns/op            23 B/op          1 allocs/op
BenchmarkTypedMapUpdate-4                        6820540               258.9  ns/op            31 B/op          2 allocs/op
BenchmarkSyncMapSimulateUpdateRange-4            7787750               201.4  ns/op            23 B/op          1 allocs/op
BenchmarkTypedMapUpdateRange-4                   7015873               260.6  ns/op            31 B/op          2 allocs/op
PASS
ok      github.com/thetechpanda/typedmap        190.395s
```

## Installation

```bash
go get github.com/thetechpanda/typedmap
```

## Contributing

Contributions are welcome and very much appreciated! 

Feel free to open an issue or submit a pull request.

## License

`TypedMap` is released under the MIT License. See the [LICENSE](LICENSE) file for details.
