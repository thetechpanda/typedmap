# typedmap
--
    import "."

typedmap package implements a simple thread-safe map that enhances sync.Map by
adding type safety and maintaining a count of the items in the map.

While TypedMap provides some similar functionality, it is not a drop-in
replacement for sync.Map and offers a simpler, more specialized interface. Its
main objective is to sacrifice some compute time for type safety and a
consistent interface.

The Keys and Values functions return slices of the keys and values in the map,
respectively. However, the order of these elements is not guaranteed to be
consistent. If the order of keys and values is important, consider using the
Entries function, which iterates over the map in a consistent order.

The Entries, Keys and Values functions do not lock the map during their
operation. They use the Len function to allocate memory for the result slices
and then use the Range function to populate these slices. Consequently, the map
could be modified by other goroutines between the time Len is called and the
Range function starts iterating, potentially leading to discrepancies.

Key, Values and Entries have the same guarantees as sync.Map.Range. When the
guarantee that the map will not be modified during the iteration is required,
use the AtomicRange function.

## Usage

#### type TypedMap

```go
type TypedMap[K any, V any] struct {
}
```

TypedMap is a thread-safe map that enhances sync.Map by adding type safety and
maintaining a count of the items in the map.

#### func  New

```go
func New[K comparable, V any]() TypedMap[K, V]
```
New returns a new TypedMap.

#### func (*TypedMap[K, V]) Clear

```go
func (m *TypedMap[K, V]) Clear()
```
Clear removes all items from the map. This is a locking operation.

#### func (*TypedMap[K, V]) Delete

```go
func (m *TypedMap[K, V]) Delete(key K)
```
Delete removes the key from the map. This is a locking operation.

#### func (*TypedMap[K, V]) Entries

```go
func (m *TypedMap[K, V]) Entries() (keys []K, values []V)
```
Entries returns two slices, one containing all the keys and the other containing
all the values present in the map.

#### func (*TypedMap[K, V]) Get

```go
func (m *TypedMap[K, V]) Get(key K) (v V, ok bool)
```
Get returns the value associated with the key and true if the key is present in
the map.

#### func (*TypedMap[K, V]) Has

```go
func (m *TypedMap[K, V]) Has(key K) bool
```
Has returns true if the map contains the key.

#### func (*TypedMap[K, V]) Keys

```go
func (m *TypedMap[K, V]) Keys() (keys []K)
```
Keys returns a slice of all the keys present in the map, an empty slice is
returned if the map is empty.

#### func (*TypedMap[K, V]) Len

```go
func (m *TypedMap[K, V]) Len() (n int)
```
Len returns the number of items in the map.

#### func (*TypedMap[K, V]) Range

```go
func (m *TypedMap[K, V]) Range(f func(K, V) bool)
```
Range calls f sequentially for each key and value present in the map. If f
returns false, Range stops the iteration.

#### func (*TypedMap[K, V]) Set

```go
func (m *TypedMap[K, V]) Set(key K, value V)
```
Set stores the key-value pair in the map. It overwrites the previous value if
the key already exists in the map. When it is important to know the previous
value, use the Update function. This is a locking operation.

#### func (*TypedMap[K, V]) Update

```go
func (m *TypedMap[K, V]) Update(key K, f func(V, bool) V)
```
Update allows the caller to change the value associated with the key atomically
guaranteeing that the value would not be changed by another goroutine during the
operation. This is a locking operation. Calling Set, Delete, Clear or Update in
f will cause a deadlock.

#### func (*TypedMap[K, V]) UpdateRange

```go
func (m *TypedMap[K, V]) UpdateRange(f func(K, V) (V, bool))
```
UpdateRange is a thread-safe version of Range that locks the map for the
duration of the iteration and allows for the modification of the values. If f
returns false, UpdateRange stops the iteration, without updating the
corresponding value in the map. Calling Set, Delete, Clear or Update within f
will cause a deadlock.

#### func (*TypedMap[K, V]) Values

```go
func (m *TypedMap[K, V]) Values() (values []V)
```
Values returns a slice of all the values present in the map, an empty slice is
returned if the map is empty.

--
**godocdown** http://github.com/robertkrimen/godocdown
