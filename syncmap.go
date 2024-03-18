package typedmap

import "github.com/thetechpanda/typedmap/internal/syncmap"

// SyncMap is a generic interface that provides a way to interact with the map.
// its just a generic wrapper around sync.Map
type SyncMap[K, V any] interface {
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
	// if the value stored in the map is equal to old.
	// The old value must be of a comparable type.
	//
	// Returns true if the swap was performed.
	CompareAndSwap(key K, old, new V) bool
	// CompareAndDelete deletes the entry for key if its value is equal to old.
	// The old value must be of a comparable type.
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

// NewSyncMap a new SyncMap that wraps sync.Map with generics.
// It allows the use of sync.Map natively and has the same drawbacks as sync.Map.
// Using CompareAndSwap or CompareAndDelete with non comparable V types will panic, as it does in sync.Map.
func NewSyncMap[K any, V any]() SyncMap[K, V] {
	return syncmap.New[K, V]()
}
