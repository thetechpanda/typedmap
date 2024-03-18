package mutex

import "reflect"

// Store sets the value for a key.
func (m *TypedMap[K, V]) Store(key K, value V) {
	m.Swap(key, value)
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *TypedMap[K, V]) Load(key K) (v V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok = m.data[key]
	return v, ok
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *TypedMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	actual, loaded = m.data[key]
	if loaded {
		return actual, true
	}
	m.data[key] = value
	return value, false
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *TypedMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	value, loaded = m.data[key]
	if loaded {
		delete(m.data, key)
	}
	return value, loaded
}

// Delete removes the key from the map.
// This is a locking operation.
func (m *TypedMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *TypedMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	previous, loaded = m.data[key]
	m.data[key] = value
	return previous, loaded
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
//
// The old value must be of a comparable type or this function will return false.
//
// Returns true if the swap was performed.
//
// ! this function uses reflect.DeepEqual to compare the values.
func (m *TypedMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	if !m.valueComparable {
		return false
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[key]
	if !ok || !reflect.DeepEqual(v, old) {
		return false
	}
	m.data[key] = new
	return true
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type or this function will return false.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
//
// ! this function uses reflect.DeepEqual to compare the values.
func (m *TypedMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	if !m.valueComparable {
		return false
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	v, ok := m.data[key]
	if !ok || !reflect.DeepEqual(v, old) {
		return false
	}
	delete(m.data, key)
	return true
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, Range stops the iteration.
// Avoid invoking any map functions within 'f' to prevent a deadlock.
func (m *TypedMap[K, V]) Range(f func(K, V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for key, value := range m.data {
		if !f(key, value) {
			break
		}
	}
}
