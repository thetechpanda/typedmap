package typedmap

import (
	"reflect"
)

// Map is a generic interface that provides a way to interact with the map.
// its interface is identical to sync.Map and so are function definition and behaviour.
type Map[K comparable, V any] interface {
	// Load returns the value stored in the map for a key, or nil if no
	// value is present.
	// The ok result indicates whether value was found in the map.
	Load(key K) (v V, ok bool)
	// Store sets the value for a key.
	Store(key K, value V)
	// LoadOrStore returns the existing value for the key if present.
	// Otherwise, it stores and returns the given value.
	// The loaded result is true if the value was loaded, false if stored.
	LoadOrStore(key K, value V) (actual V, loaded bool)
	// LoadAndDelete deletes the value for a key, returning the previous value if any.
	// The loaded result reports whether the key was present.
	LoadAndDelete(key K) (value V, loaded bool)
	// Delete deletes the value for a key.
	Delete(key K)
	// Swap swaps the value for a key and returns the previous value if any.
	// The loaded result reports whether the key was present.
	Swap(key K, value V) (previous V, loaded bool)
	// CompareAndSwap swaps the old and new values for key
	// if V is not comparable type this function will return false.
	// ! this function uses reflect.DeepEqual to compare the values.
	//
	// if the value stored in the map is equal to old.
	CompareAndSwap(key K, old, new V) bool
	// CompareAndDelete deletes the entry for key if its value is equal to old.
	// if V is not comparable type this function will return false.
	// ! this function uses reflect.DeepEqual to compare the values.
	//
	// If there is no current value for key in the map, CompareAndDelete
	// returns false (even if the old value is the nil interface value).
	CompareAndDelete(key K, old V) (deleted bool)
	// Range calls f sequentially for each key and value present in the map.
	// If f returns false, range stops the iteration.
	//
	// Range does not necessarily correspond to any consistent snapshot of the Map's
	// contents: no key will be visited more than once, but if the value for any key
	// is stored or deleted concurrently (including by f), Range may reflect any
	// mapping for that key from any point during the Range call. Range does not
	// block other methods on the receiver; even f itself may call any method on m.
	//
	// Range may be O(N) with the number of elements in the map even if f returns
	// false after a constant number of calls.
	Range(f func(K, V) bool)
}

// Store sets the value for a key.
func (m *typedMap[K, V]) Store(key K, value V) {
	m.Swap(key, value)
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *typedMap[K, V]) Load(key K) (v V, ok bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok = m.data[key]
	return v, ok
}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *typedMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
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
func (m *typedMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
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
func (m *typedMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *typedMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	previous, loaded = m.data[key]
	m.data[key] = value
	return previous, loaded
}

// CompareAndSwap swaps the old and new values for key
// if V is not comparable type this function will return false.
// ! this function uses reflect.DeepEqual to compare the values.
//
// if the value stored in the map is equal to old.
func (m *typedMap[K, V]) CompareAndSwap(key K, old, new V) bool {
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
// if V is not comparable type this function will return false.
// ! this function uses reflect.DeepEqual to compare the values.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (m *typedMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
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
func (m *typedMap[K, V]) Range(f func(K, V) bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	for key, value := range m.data {
		if !f(key, value) {
			break
		}
	}
}
