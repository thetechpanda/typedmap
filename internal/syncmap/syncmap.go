package syncmap

import (
	"sync"
)

type SyncMap[K, V any] struct {
	sm sync.Map
}

func New[K, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{}
}

// Store sets the value for a key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.sm.Store(key, value)
}

// Load returns the value stored in the map for a key, or nil if no
// value is present.
// The ok result indicates whether value was found in the map.
func (m *SyncMap[K, V]) Load(key K) (v V, ok bool) {
	value, ok := m.sm.Load(key)
	if value == nil {
		return m.zeroValue(), ok
	}
	return value.(V), ok

}

// LoadOrStore returns the existing value for the key if present.
// Otherwise, it stores and returns the given value.
// The loaded result is true if the value was loaded, false if stored.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	v, loaded := m.sm.LoadOrStore(key, value)
	if v == nil {
		return m.zeroValue(), loaded
	}
	return v.(V), loaded
}
func (m *SyncMap[K, V]) zeroValue() V {
	var zero V
	return zero
}

// LoadAndDelete deletes the value for a key, returning the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	v, loaded := m.sm.LoadAndDelete(key)
	if v == nil {
		return m.zeroValue(), loaded
	}
	return v.(V), loaded
}

// Delete removes the key from the map.
// This is a locking operation.
func (m *SyncMap[K, V]) Delete(key K) {
	m.LoadAndDelete(key)
}

// Swap swaps the value for a key and returns the previous value if any.
// The loaded result reports whether the key was present.
func (m *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	previousAny, loaded := m.sm.Swap(key, value)
	if previousAny == nil {
		return m.zeroValue(), loaded
	}
	return previousAny.(V), loaded
}

// CompareAndSwap swaps the old and new values for key
// if the value stored in the map is equal to old.
// The old value must be of a comparable type.
//
// Returns true if the swap was performed.
func (m *SyncMap[K, V]) CompareAndSwap(key K, old, new V) bool {
	return m.sm.CompareAndSwap(key, old, new)
}

// CompareAndDelete deletes the entry for key if its value is equal to old.
// The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete
// returns false (even if the old value is the nil interface value).
func (m *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return m.sm.CompareAndDelete(key, old)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, Range stops the iteration.
// Avoid invoking any map functions within 'f' to prevent a deadlock.
func (m *SyncMap[K, V]) Range(f func(K, V) bool) {
	m.sm.Range(func(key, value any) bool {
		var k K
		var v V
		if key != nil {
			k = key.(K)
		}
		if value != nil {
			v = value.(V)
		}
		return f(k, v)
	})
}
