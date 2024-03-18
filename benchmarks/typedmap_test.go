package benchmarks

import (
	"testing"

	"github.com/thetechpanda/typedmap"
)

func BenchmarkTypedMapStoreAndDelete(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Delete(i)
	}
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

func BenchmarkTypedMapLoad(b *testing.B) {
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

func BenchmarkTypedMapEntries(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	keys, values := m.Entries()
	noop(keys, values)
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

func BenchmarkTypedMapValues(b *testing.B) {
	m := typedmap.New[int, int]()
	for i := 0; i < b.N; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()
	values := m.Values()
	noop(values)
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

func BenchmarkTypedMapConcurrentOperations(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
		m.Load(i)
		m.Delete(i)
	})
}

func BenchmarkTypedMapConcurrentStore(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Store(i, j)
	})
}

func BenchmarkTypedMapConcurrentSwap(b *testing.B) {
	m := typedmap.New[int, int]()
	benchmarkConcurrentInt(b, func(n, i, j int) {
		m.Swap(i, j)
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
