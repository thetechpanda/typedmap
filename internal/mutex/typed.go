package mutex

import "sync"

// TypedMap implements a simple thread-safe map that uses generics.
type TypedMap[K comparable, V any] struct {
	mu              sync.RWMutex
	valueComparable bool
	data            map[K]V
}

// Clear removes all items from the map.
// This is a locking operation.
func (m *TypedMap[K, V]) Clear() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data = make(map[K]V)
}

// Has returns true if the map contains the key.
func (m *TypedMap[K, V]) Has(key K) bool {
	_, ok := m.Load(key)
	return ok
}

// Update allows the caller to change the value associated with the key atomically guaranteeing that the value would not be changed by another goroutine during the operation.
//
// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
func (m *TypedMap[K, V]) Update(key K, f func(V, bool) V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[key]
	m.data[key] = f(v, ok)
}

// UpdateRange is a thread-safe version of Range that locks the map for the duration of the iteration and allows for the modification of the values.
// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the map.
//
// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
func (m *TypedMap[K, V]) UpdateRange(f func(K, V) (V, bool)) {
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
//
// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
func (m *TypedMap[K, V]) Exclusive(f func(m map[K]V)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	f(m.data)
}

// Len returns the number of items in the map.
func (m *TypedMap[K, V]) Len() (n int) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.data)
}

// Keys returns a slice of all the keys present in the map, an empty slice is returned if the map is empty.
func (m *TypedMap[K, V]) Keys() (keys []K) {
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
func (m *TypedMap[K, V]) Values() (values []V) {
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
func (m *TypedMap[K, V]) Entries() (keys []K, values []V) {
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
