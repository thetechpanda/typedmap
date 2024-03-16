// The benchmarks are designed to compare the performance of TypedMap with sync.Map, since sync.Map is not typed, the benchmarks are designed to compare the performance of the operations that are common between the two maps.
package typedmap_test

import (
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap"
)

// use is an helper function that does nothing with the input.
func noop(x ...any) {
	_ = x
}

// BenchmarkConcurrentSyncMapStore measures the performance of concurrent storing of values in sync.Map.
func BenchmarkConcurrentSyncMapStore(b *testing.B) {
	var m sync.Map
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Store(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkConcurrentTypedMapSet measures the performance of concurrent setting of values in TypedMap.
func BenchmarkConcurrentTypedMapSet(b *testing.B) {
	m := typedmap.New[int, int]()
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(i, i)
		}(i)
	}
	wg.Wait()
}

// BenchmarkSyncMapDelete measures the performance of deleting values from sync.Map.
func BenchmarkSyncMapDelete(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

// BenchmarkTypedMapDelete measures the performance of deleting values from TypedMap.
func BenchmarkTypedMapDelete(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

// BenchmarkSyncMapRange measures the performance of iterating over sync.Map.
func BenchmarkSyncMapRange(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	m.Range(func(k, v any) bool {
		kk, vv := k.(int), v.(int)
		noop(kk, vv)
		return true
	})
}

// BenchmarkTypedMapRange measures the performance of iterating over TypedMap.
func BenchmarkTypedMapRange(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	m.Range(func(k int, v int) bool {
		noop(k, v)
		return true
	})
}

// BenchmarkSyncMapRange measures the performance of iterating over sync.Map.
func BenchmarkSyncMapLoad(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _ := m.Load(i)
		noop(v.(int))
	}
}

// BenchmarkTypedMapRange measures the performance of iterating over TypedMap.
func BenchmarkTypedMapGet(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _ := m.Get(i)
		noop(v)
	}
}

// BenchmarkSyncMapSimulateEntries measures the performance of returning all keys and values from the map.
// as there is no direct way to get all keys and values from sync.Map, we simulate it by iterating over all keys and getting their values.
func BenchmarkSyncMapSimulateEntries(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys, values := make([]any, 0), make([]any, 0)
	m.Range(func(k, v any) bool {
		keys = append(keys, k)
		values = append(values, v)
		return true
	})
	noop(keys, values)
}

// BenchmarkTypedMapEntries measures the performance of Entries.
func BenchmarkTypedMapEntries(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	keys, values := m.Entries()
	noop(keys, values)
}

// BenchmarkSyncMapSimulateKeys measures the performance of returning all keys from the map.
// As there is no direct way to get all keys from sync.Map, we simulate it by iterating over all keys using range.
func BenchmarkSyncMapSimulateKeys(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys := make([]any, 0)
	m.Range(func(k, v any) bool {
		keys = append(keys, k)
		return true
	})
	noop(keys)
}

// BenchmarkTypedMapKeys measures the performance of Keys.
func BenchmarkTypedMapKeys(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	keys := m.Keys()
	noop(keys)
}

// BenchmarkSyncMapSimulateKeys measures the performance of returning all values from the map.
// As there is no direct way to get all values from sync.Map, we simulate it by iterating over all keys and getting their values.
func BenchmarkSyncMapSimulateValues(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	values := make([]any, 0)
	m.Range(func(k, v any) bool {
		values = append(values, v.(int))
		return true
	})
	noop(values)
}

// BenchmarkTypedMapValues measures the performance of Values.
func BenchmarkTypedMapValues(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	values := m.Values()
	noop(values)
}

// BenchmarkSyncMapSimulateUpdate measures the performance using a mutex to simulate updating values in sync.Map assuming that the mutex is shared with other operations.
func BenchmarkSyncMapSimulateUpdate(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < b.N; i++ {
		k, _ := m.Load(i)
		m.Store(k, k.(int)*i)
	}

}

// BenchmarkTypedMapUpdate measures the performance of Update.
func BenchmarkTypedMapUpdate(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(i, func(v int, ok bool) int {
			return v * i
		})

	}
}

// BenchmarkSyncMapSimulateUpdateRange measures the performance using a mutex to simulate updating values in sync.Map assuming that the mutex is shared with other operations.
func BenchmarkSyncMapSimulateUpdateRange(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	m.Range(func(k, v any) bool {
		noop(k.(int), v.(int))
		m.Store(k, v.(int)+1)
		return true
	})

}

// BenchmarkTypedMapAtomicRange measures the performance of iterating over TypedMap using AtomicRange.
func BenchmarkTypedMapUpdateRange(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()
	m.UpdateRange(func(k, i int) (int, bool) {
		noop(k, i)
		return i + 1, true
	})
}
