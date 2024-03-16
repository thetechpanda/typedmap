// typedmap package implements a simple thread-safe map that enhances sync.Map by adding type safety and maintaining a count of the items in the map.
//
// While TypedMap provides some similar functionality, it is not a drop-in replacement for sync.Map and offers a simpler, more specialized interface. Its main objective is to sacrifice some compute time for type safety and a consistent interface.
//
// The Keys and Values functions return slices of the keys and values in the map, respectively. However, the order of these elements is not guaranteed to be consistent.
// If the order of keys and values is important, consider using the Entries function, which iterates over the map in a consistent order.
//
// The Entries, Keys and Values functions do not lock the map during their operation. They use the Len function to allocate memory for the result slices and then use the Range function to populate these slices.
// Consequently, the map could be modified by other goroutines between the time Len is called and the Range function starts iterating, potentially leading to discrepancies.
//
// Key, Values and Entries have the same guarantees as sync.Map.Range. When the guarantee that the map will not be modified during the iteration is required, use the AtomicRange function.
package typedmap

import (
	"sync"
	"sync/atomic"
)

// TypedMap is a thread-safe map that enhances sync.Map by adding type safety and maintaining a count of the items in the map.
type TypedMap[K any, V any] struct {
	mu    sync.Mutex
	count atomic.Int64
	data  sync.Map
}

// New returns a new TypedMap.
func New[K comparable, V any]() TypedMap[K, V] {
	return TypedMap[K, V]{}
}

// Get returns the value associated with the key and true if the key is present in the map.
func (m *TypedMap[K, V]) Get(key K) (v V, ok bool) {
	x, ok := m.data.Load(key)
	if !ok {
		return v, false
	}
	return x.(V), true
}

// Set stores the key-value pair in the map. It overwrites the previous value if the key already exists in the map.
// When it is important to know the previous value, use the Update function.
// This is a locking operation.
func (m *TypedMap[K, V]) Set(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.unsafeSet(key, value)
}

// unsafeSet is a non-locking version of Set. It is used by Set and Update to avoid deadlocks.
func (m *TypedMap[K, V]) unsafeSet(key K, value V) {
	if !m.Has(key) {
		m.count.Add(1)
	}
	m.data.Store(key, value)
}

// Update allows the caller to change the value associated with the key atomically guaranteeing that the value would not be changed by another goroutine during the operation.
// This is a locking operation. Calling Set, Delete, Clear or Update in f will cause a deadlock.
func (m *TypedMap[K, V]) Update(key K, f func(V, bool) V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.Get(key)
	m.unsafeSet(key, f(v, ok))
}

// Clear removes all items from the map.
// This is a locking operation.
func (m *TypedMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = sync.Map{}
	m.count.Store(0)
}

// Delete removes the key from the map.
// This is a locking operation.
func (m *TypedMap[K, V]) Delete(key K) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.Has(key) {
		m.data.Delete(key)
		m.count.Add(-1)
	}
}

// Has returns true if the map contains the key.
func (m *TypedMap[K, V]) Has(key K) bool {
	_, ok := m.data.Load(key)
	return ok
}

// Keys returns a slice of all the keys present in the map, an empty slice is returned if the map is empty.
func (m *TypedMap[K, V]) Keys() (keys []K) {
	// map could have been modified between the time we got the count and the time we started the range
	// max ensures two things 1. we don't allocate more memory than we need 2. we don't iterate more than we need
	max := m.Len()
	if max == 0 {
		return make([]K, 0)
	}
	keys = make([]K, 0, max)
	m.data.Range(func(key, value any) (ok bool) {
		keys = append(keys, key.(K))
		return len(keys) <= max
	})
	return keys
}

// Values returns a slice of all the values present in the map, an empty slice is returned if the map is empty.
func (m *TypedMap[K, V]) Values() (values []V) {
	// map could have been modified between the time we got the count and the time we started the range
	// max ensures two things 1. we don't allocate more memory than we need 2. we don't iterate more than we need
	max := m.Len()
	if max == 0 {
		return make([]V, 0)
	}
	values = make([]V, 0, max)
	m.data.Range(func(key, value any) bool {
		values = append(values, value.(V))
		return len(values) <= max

	})
	return values
}

// Entries returns two slices, one containing all the keys and the other containing all the values present in the map.
func (m *TypedMap[K, V]) Entries() (keys []K, values []V) {
	max := m.Len()
	if max == 0 {
		return make([]K, 0), make([]V, 0)
	}
	keys = make([]K, 0, max)
	values = make([]V, 0, max)
	m.data.Range(func(key, value any) (ok bool) {
		keys = append(keys, key.(K))
		values = append(values, value.(V))
		return len(keys) <= max
	})
	return keys, values
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, Range stops the iteration.
func (m *TypedMap[K, V]) Range(f func(K, V) bool) {
	m.data.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

// UpdateRange is a thread-safe version of Range that locks the map for the duration of the iteration and allows for the modification of the values.
// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the map.
// Calling Set, Delete, Clear or Update within f will cause a deadlock.
func (m *TypedMap[K, V]) UpdateRange(f func(K, V) (V, bool)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data.Range(func(key, value any) bool {
		v, ok := f(key.(K), value.(V))
		if !ok {
			return false
		}
		m.unsafeSet(key.(K), v)
		return true
	})
}

// Len returns the number of items in the map.
func (m *TypedMap[K, V]) Len() (n int) {
	return int(m.count.Load())
}
