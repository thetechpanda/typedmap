package benchmarks

import (
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap"
)

func BenchmarkTypedSyncMapStoreAndDelete(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

func BenchmarkTypedSyncMapRange(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	m.Range(func(k, v int) bool {
		noop(k, v)
		return true
	})
}

func BenchmarkTypedSyncMapLoad(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _ := m.Load(i)
		noop(v)
	}
}

func BenchmarkTypedSyncMapSimulateEntries(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys, values := make([]int, 0), make([]int, 0)
	m.Range(func(k, v int) bool {
		keys = append(keys, k)
		values = append(values, v)
		return true
	})
	noop(keys, values)
}

func BenchmarkTypedSyncMapSimulateKeys(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys := make([]int, 0)
	m.Range(func(k, v int) bool {
		keys = append(keys, k)
		return true
	})
	noop(keys)
}

func BenchmarkTypedSyncMapSimulateValues(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	values := make([]int, 0)
	m.Range(func(k, v int) bool {
		values = append(values, v)
		return true
	})
	noop(values)
}

func BenchmarkTypedSyncMapSimulateUpdate(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	for i := 0; i < b.N; i++ {
		k, _ := m.Load(i)
		m.Store(k, k*i)
	}

}

func BenchmarkTypedSyncMapSimulateUpdateRange(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	mu := sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	m.Range(func(k, v int) bool {
		noop(k, v)
		m.Store(k, v+1)
		return true
	})

}

func BenchmarkTypedSyncMapConcurrentOperations(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
		m.Load(i)
		m.Delete(j)
	})
}

func BenchmarkTypedSyncMapConcurrentStore(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
	})
}

func BenchmarkTypedSyncMapConcurrentSwap(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Swap(i, j)
	})
}

func BenchmarkTypedSyncMapConcurrentLoadOrStore(b *testing.B) {
	m := typedmap.NewSyncMap[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		_, ok := m.LoadOrStore(i, j)
		if !ok {
			m.Delete(i)
		}
	})
}
