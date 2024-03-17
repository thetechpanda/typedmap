// typedmap package implements a simple thread-safe map that uses generics.
//
// The Keys and Values functions return slices of the keys and values in the map, respectively. However, the order of these elements is not guaranteed to be consistent.
// If the order of keys and values is important, consider using the Entries function, which iterates over the map in a consistent order.
//
// The read operations use RWMutex.RLock to allow multiple readers to access the map concurrently, while the write operations use RWMutex.Lock to ensure that only one writer can access the map at a time.
//
// There are a few things to note about the implementation when using the sync.Map interface:
//
//   - sync.Map uses K, V any, which means that the keys and values can be of any type. However, the typedmap package uses K comparable, V any, which means that the keys must be comparable.
//   - The CompareAndSwap and CompareAndDelete functions use reflect.DeepEqual to compare the values, which may not be as efficient as using the == operator for simple types. TypeMap detects if the value is comparable type and will always return false if it is not.
package typedmap

import (
	"reflect"
	"sync"
)

// TypedMap is a generic interface that provides a way to interact with the map.
// Its interface extends Map[K, V]
type TypedMap[K comparable, V any] interface {
	Map[K, V]
	// Update allows the caller to change the value associated with the key atomically guaranteeing that the value would not be changed by another goroutine during the operation.
	// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
	Update(key K, f func(V, bool) V)
	// UpdateRange is a thread-safe version of Range that locks the map for the duration of the iteration and allows for the modification of the values.
	// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the map.
	// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
	UpdateRange(f func(K, V) (V, bool))
	// Exclusive provides a way to perform  operations on the map ensuring that no other operation is performed on the map during the execution of the function.
	// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
	Exclusive(f func(m map[K]V))
	// Clear removes all items from the map.
	Clear()
	// Has returns true if the map contains the key.
	Has(key K) bool
	// Keys returns a slice of all the keys present in the map, an empty slice is returned if the map is empty.
	Keys() (keys []K)
	// Values returns a slice of all the values present in the map, an empty slice is returned if the map is empty.
	Values() (values []V)
	// Entries returns two slices, one containing all the keys and the other containing all the values present in the map.
	Entries() (keys []K, values []V)
	// Len returns the number of unique keys in the map.
	Len() (n int)
}

// TypedMap implements a simple thread-safe map that uses generics.
type typedMap[K comparable, V any] struct {
	mu              sync.RWMutex
	valueComparable bool
	data            map[K]V
}

// New returns a new TypedMap.
func New[K comparable, V any]() TypedMap[K, V] {
	return NewWithMap[K, V](nil)
}

// NewWithMap returns a new TypedMap, initialized with the given map. if m is nil, an empty map is created.
// m key, values are copied, so that the caller can safely modify the map after creating a TypedMap.
func NewWithMap[K comparable, V any](m map[K]V) TypedMap[K, V] {
	var v map[K]V = make(map[K]V, len(m))
	for key, value := range m {
		v[key] = value
	}
	var z V
	return &typedMap[K, V]{data: v, valueComparable: reflect.TypeOf(z).Comparable()}
}

// NewSyncMapCompatible returns a new TypedMap that is compatible with sync.Map interface.
// You can always cast TypeMap to Map[K, V] and use it as you would sync.Map with the added benefit of type safety.
func NewSyncMapCompatible[K comparable, V any]() Map[K, V] {
	return NewWithMap[K, V](nil)
}

// Clear removes all items from the map.
// This is a locking operation.
func (m *typedMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[K]V)
}

// Has returns true if the map contains the key.
func (m *typedMap[K, V]) Has(key K) bool {
	_, ok := m.Load(key)
	return ok
}

// Update allows the caller to change the value associated with the key atomically guaranteeing that the value would not be changed by another goroutine during the operation.
// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
func (m *typedMap[K, V]) Update(key K, f func(V, bool) V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[key]
	m.data[key] = f(v, ok)
}

// UpdateRange is a thread-safe version of Range that locks the map for the duration of the iteration and allows for the modification of the values.
// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the map.
// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
func (m *typedMap[K, V]) UpdateRange(f func(K, V) (V, bool)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	for key, value := range m.data {
		newValue, ok := f(key, value)
		if !ok {
			return
		}
		m.data[key] = newValue
	}
}

// Exclusive provides a way to perform  operations on the map ensuring that no other operation is performed on the map during the execution of the function.
// ! Avoid invoking any map functions within 'f' to prevent a deadlock.
func (m *typedMap[K, V]) Exclusive(f func(m map[K]V)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// Len returns the number of items in the map.
func (m *typedMap[K, V]) Len() (n int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Keys returns a slice of all the keys present in the map, an empty slice is returned if the map is empty.
func (m *typedMap[K, V]) Keys() (keys []K) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	max := len(m.data)
	if max == 0 {
		return make([]K, 0)
	}
	keys = make([]K, 0, max)
	for key := range m.data {
		keys = append(keys, key)
		if len(keys) >= max {
			break
		}
	}
	return keys
}

// Values returns a slice of all the values present in the map, an empty slice is returned if the map is empty.
func (m *typedMap[K, V]) Values() (values []V) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	max := len(m.data)
	if max == 0 {
		return make([]V, 0)
	}
	values = make([]V, 0, max)
	for _, value := range m.data {
		values = append(values, value)
		if len(values) >= max {
			break
		}
	}
	return values
}

// Entries returns two slices, one containing all the keys and the other containing all the values present in the map.
func (m *typedMap[K, V]) Entries() (keys []K, values []V) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	max := len(m.data)
	if max == 0 {
		return make([]K, 0), make([]V, 0)
	}
	keys = make([]K, 0, max)
	values = make([]V, 0, max)
	for key, value := range m.data {
		keys = append(keys, key)
		values = append(values, value)
		if len(keys) >= max {
			break
		}
	}
	return keys, values
}
