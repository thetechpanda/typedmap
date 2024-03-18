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

import "github.com/thetechpanda/typedmap/internal/mutex"

// TypedMap is a generic interface that provides a way to interact with the map.
// Its interface extends Map[K, V]
type TypedMap[K comparable, V any] interface {
	Map[K, V]
	// Update allows the caller to change the value associated with the key atomically guaranteeing that the value would not be changed by another goroutine during the operation.
	//
	// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
	Update(key K, f func(V, bool) V)
	// UpdateRange is a thread-safe version of Range that locks the map for the duration of the iteration and allows for the modification of the values.
	// If f returns false, UpdateRange stops the iteration, without updating the corresponding value in the map.
	//
	// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
	UpdateRange(f func(K, V) (V, bool))
	// Exclusive provides a way to perform  operations on the map ensuring that no other operation is performed on the map during the execution of the function.
	//
	// ! Do not invoke any TypedMap functions within 'f' to prevent a deadlock.
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

// New returns a new TypedMap.
func New[K comparable, V any]() TypedMap[K, V] {
	return mutex.New(map[K]V{})
}

// NewWithMap returns a new TypedMap, initialized with the given map. if m is nil, an empty map is created.
// m key, values are copied, so that the caller can safely modify the map after creating a TypedMap.
func NewWithMap[K comparable, V any](m map[K]V) TypedMap[K, V] {
	return mutex.New(m)
}
