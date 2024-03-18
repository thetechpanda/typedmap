package benchmarks

import (
	"sync"
	"testing"
)

func BenchmarkNativeSyncMapStoreAndDelete(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

func BenchmarkNativeSyncMapRange(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	m.Range(func(k, v any) bool {
		noop(k.(int), v.(int))
		return true
	})
}

func BenchmarkNativeSyncMapLoad(b *testing.B) {
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

func BenchmarkNativeSyncMapSimulateEntries(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys, values := make([]int, 0), make([]int, 0)
	m.Range(func(k, v any) bool {
		keys = append(keys, k.(int))
		values = append(values, v.(int))
		return true
	})
	noop(keys, values)
}

func BenchmarkNativeSyncMapSimulateKeys(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys := make([]int, 0)
	m.Range(func(k, v any) bool {
		keys = append(keys, k.(int))
		return true
	})
	noop(keys)
}

func BenchmarkNativeSyncMapSimulateValues(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	values := make([]int, 0)
	m.Range(func(k, v any) bool {
		values = append(values, v.(int))
		return true
	})
	noop(values)
}

func BenchmarkNativeSyncMapSimulateUpdate(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < b.N; i++ {
		v, _ := m.Load(i)
		m.Store(i, v.(int)*i)
	}

}

func BenchmarkNativeSyncMapSimulateUpdateRange(b *testing.B) {
	var m sync.Map
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	m.Range(func(k, v any) bool {
		noop(k, v)
		m.Store(k.(int), v.(int)+1)
		return true
	})

}

func BenchmarkNativeSyncMapConcurrentOperations(b *testing.B) {
	var m sync.Map
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
		v, _ := m.Load(i)
		if v != nil {
			_ = v.(int)
		}
		m.Delete(j)
	})
}

func BenchmarkNativeSyncMapConcurrentStore(b *testing.B) {
	var m sync.Map
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
	})
}

func BenchmarkNativeSyncMapConcurrentSwap(b *testing.B) {
	var m sync.Map
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Swap(i, j)
	})
}

func BenchmarkNativeSyncMapConcurrentLoadOrStore(b *testing.B) {
	var m sync.Map
	benchmarkConcurrentInt(b, func(n, i, j int) {
		v, ok := m.LoadOrStore(i, j)
		if !ok {
			m.Delete(i)
		} else if v != nil {
			_ = v.(int)
		}
	})
}
