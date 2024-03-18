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
