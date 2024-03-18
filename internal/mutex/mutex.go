package mutex

import "reflect"

// New returns a new TypedMap, initialized with the given map. if m is nil, an empty map is created.
// m key, values are copied, so that the caller can safely modify the map after creating a TypedMap.
func New[K comparable, V any](m map[K]V) *TypedMap[K, V] {
	var v map[K]V = make(map[K]V, len(m))
	for key, value := range m {
		v[key] = value
	}
	var z V
	return &TypedMap[K, V]{data: v, valueComparable: reflect.TypeOf(z).Comparable()}
}
