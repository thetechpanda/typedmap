package typedmap_test

import (
	"context"
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap"
)

// use is an helper function that does nothing with the input.
func noop(x ...any) {
	_ = x
}

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

func BenchmarkConcurrentTypedMapSet(b *testing.B) {
	m := typedmap.New[int, int]()
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

func BenchmarkTypedMapDelete(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
}

func BenchmarkSyncMapRange(b *testing.B) {
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

func BenchmarkTypedMapRange(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	m.Range(func(k int, v int) bool {
		noop(k, v)
		return true
	})
}

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

func BenchmarkTypedMapGet(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v, _ := m.Load(i)
		noop(v)
	}
}

func BenchmarkSyncMapSimulateEntries(b *testing.B) {
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

func BenchmarkTypedMapEntries(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys, values := m.Entries()
	noop(keys, values)
}

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

func BenchmarkTypedMapKeys(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys := m.Keys()
	noop(keys)
}

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

func BenchmarkTypedMapValues(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	values := m.Values()
	noop(values)
}

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

func BenchmarkTypedMapUpdate(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Update(i, func(v int, ok bool) int {
			return v * i
		})

	}
}

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

func BenchmarkTypedMapUpdateRange(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	m.UpdateRange(func(k, i int) (int, bool) {
		noop(k, i)
		return i + 1, true
	})
}

func benchmarkConcurrentInt(b *testing.B, f func(n, i, j int)) {
	numGoroutines := 100
	numOperations := 1000
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func() {
				<-ctx.Done()
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					f(n, j, i)
				}
			}()
		}
	}
	cancel()
	wg.Wait()
}

func BenchmarkSyncMapConcurrentOperations(b *testing.B) {
	var m = &sync.Map{}
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
		m.Load(i)
		m.Delete(j)
	})
}

func BenchmarkTypedMapConcurrentOperations(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
		m.Load(i)
		m.Delete(i)
	})
}
func BenchmarkSyncMapConcurrentStore(b *testing.B) {
	m := &sync.Map{}
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
	})
}

func BenchmarkTypedMapConcurrentStore(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
	})
}
func BenchmarkSyncMapConcurrentSwap(b *testing.B) {
	m := &sync.Map{}
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Swap(i, j)
	})
}

func BenchmarkTypedMapConcurrentSwap(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Swap(i, j)
	})
}

func BenchmarkSyncMapConcurrentLoadOrStore(b *testing.B) {
	m := &sync.Map{}
	benchmarkConcurrentInt(b, func(n, i, j int) {
		v, ok := m.LoadOrStore(i, j)
		_ = v.(int)
		if !ok {
			m.Delete(i)
		}
	})
}

func BenchmarkTypedMapConcurrentLoadOrStore(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		_, ok := m.LoadOrStore(i, j)
		if !ok {
			m.Delete(i)
		}
	})
}

func BenchmarkTypedMapConcurrentUpdate(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Update(i, func(v int, ok bool) int {
			return v + 1
		})
	})
}
