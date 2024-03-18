# TypedMap

[![Go Report Card](https://goreportcard.com/badge/github.com/thetechpanda/typedmap)](https://goreportcard.com/report/github.com/thetechpanda/typedmap)
[![Go Reference](https://pkg.go.dev/badge/github.com/thetechpanda/typedmap.svg)](https://pkg.go.dev/github.com/thetechpanda/typedmap)
[![Release](https://img.shields.io/github/release/thetechpanda/typedmap.svg?style=flat-square)](https://github.com/thetechpanda/typedmap/tags/latest)

`TypedMap` implements a simple thread-safe map that behaves similarly to `sync.Map` adding type safety and making it simple to know how many unique keys are in the map, `TypeMap` implements the same interface as `sync.Map` and provides `Map[K, V]` interface to ease code refactoring from `sync.Map` to `TypedMap`. 

While `TypedMap` provides some similar functionality to `sync.Map`, it is not a drop-in replacement as at is core it uses RWMutex, while `sync.Map` has a more specialised use case. See [sync.Map](https://pkg.go.dev/sync#Map) for its specific use cases.

If your use case falls into `sync.Map` use cases, use `SyncMap`, that wraps `sync.Map` with generics and provide the minimum amount of code to ensure typecasting does not panic. Check out the benchmarks for an idea of the overhead the changes introduce.

## Pointers Values

Consider the following when using `TypedMap` with pointer values:

* **Concurrent Modification:** If multiple goroutines modify the data pointed to by the same pointer without proper synchronization, it can lead to race conditions and unpredictable behavior.
* **Data Race:** Even if the `TypedMap` itself is thread-safe, the data pointed to by the values is not automatically protected. Accessing or modifying the data through pointers in concurrent goroutines can cause data races.

## Migrating from sync.Map to TypedMap
`TypedMap[K comparable, V any]` and `Map[K comparable, V any]` works only with `comparable` keys as `V any` cannot be used as map's keys.

`SyncMap[K, V any]` has the same interface as `sync.Map` and you could pass non-comparable types as a key, but `sync.Map` will, anyway, panic in such case, eg:

```
panic: runtime error: hash of unhashable type []int [recovered]
	panic: runtime error: hash of unhashable type []int
```

## Documentation

You can find the generated go doc [here](godoc.txt).

## Key Features

* **Type Safety:** Uses generics to provide a type-safe interface for keys and values, eliminating the need for specialised structs and interfaces.
* **Thread Safety:** Ensures safe concurrent access to the map through the use of a sync.RWMutex.
* **Atomic Updates:** Includes functions that allows for atomic modifications to values in the map.
* **Iteration:** Supports iterating over the map with the Range function, and provides methods to obtain slices of keys (Keys), values (Values), or both (Entries).
* **Map Size:** Offers a Len function to easily retrieve the number of items in the map.
* **Typed sync.Map:** `SyncMap[K, V any]` can be used as a drop in replacement for `sync.Map`, at its core uses `sync.Map` itself.

## Motivation

Go's `sync.Map` has a general-purpose design, optimized for use cases where keys are dynamically added and removed by multiple goroutines. 

However, this flexibility comes at the cost of type safety and can lead to more verbose and complex code. 

`TypedMap`, on the other hand, leverages generics to offer a more streamlined and type-safe interface, making it easier to work with while still providing the necessary thread safety for concurrent operations. `SyncMap` does the same using a native `sync.Map`.

## Usage

```go
package main

import (
	"fmt"
	"github.com/thetechpanda/typedmap"
)

func main() {
	m := typedmap.New[string, int]()
	// or using sync.Map
	m := typedmap.NewSyncMap[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	// v is int
}
```

## Benchmarks

Check [BENCHMARKS](BENCHMARKS.md) for more information.

## Installation

```bash
go get github.com/thetechpanda/typedmap
```

## Contributing

Contributions are welcome and very much appreciated! 

Feel free to open an issue or submit a pull request.

## License

`TypedMap` is released under the MIT License. See the [LICENSE](LICENSE) file for details.
