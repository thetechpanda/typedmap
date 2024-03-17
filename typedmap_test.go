package typedmap_test

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/thetechpanda/typedmap"
)

func TestNew(t *testing.T) {
	m := typedmap.New[string, int]()
	if m.Len() != 0 {
		t.Errorf("Len(): Expected a new map, got map with length %d", m.Len())
	}
}
func TestNewWithMap(t *testing.T) {
	data := map[string]int{
		"key1": 42,
		"key2": 43,
	}
	m := typedmap.NewWithMap(data)
	v, ok := m.Load("key1")
	if !ok || v != 42 {
		t.Errorf("Load(): Expected value 42 for key %q, got value %d", "key1", v)
	}
	v, ok = m.Load("key2")
	if !ok || v != 43 {
		t.Errorf("Load(): Expected value 43 for key %q, got value %d", "key2", v)
	}
}

func TestLoad(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok || v != value {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", value, key, v)
	}
}

func TestStore(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok || v != value {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", value, key, v)
	}
}

func TestLoadOrStore(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	actual, loaded := m.LoadOrStore(key, value)
	if loaded {
		t.Errorf("LoadOrStore(): Expected key %q to be stored", key)
	}
	if actual != value {
		t.Errorf("Expected value %d for key %q, got value %d", value, key, actual)
	}

	actual, loaded = m.LoadOrStore(key, 43)
	if !loaded {
		t.Errorf("LoadOrStore(): Expected key %q to be loaded", key)
	}
	if actual != value {
		t.Errorf("Expected value %d for key %q, got value %d", value, key, actual)
	}
}

func TestLoadAndDelete(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	m.Store(key, 42)
	v, deleted := m.LoadAndDelete(key)
	if !deleted {
		t.Errorf("LoadAndDelete(): Expected key %q to be deleted", key)
	}
	if _, ok := m.Load(key); ok {
		t.Errorf("Load(): Expected key %q to be deleted", key)
	} else if v != 42 {
		t.Errorf("Expected value 42, got %d", v)
	}
}

func TestDelete(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	m.Delete(key)
	if _, ok := m.Load(key); ok {
		t.Errorf("Load(): Expected key %q to be deleted", key)
	}
}

func TestSwap(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	previous, loaded := m.Swap(key, value)
	if loaded {
		t.Errorf("Swap(): Key %q was present in an empty map", key)
	}
	if previous != 0 {
		t.Errorf("Expected previous value 0, got %d", previous)
	}

	previous, loaded = m.Swap(key, 43)
	if !loaded {
		t.Errorf("Swap(): Expected key %q to be loaded", key)
	}
	if previous != value {
		t.Errorf("Expected previous value %d, got %d", value, previous)
	}
}

func TestCompareAndSwap(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	current, swap := 42, 43
	m.Store(key, current)
	swapped := m.CompareAndSwap(key, current, swap)
	if !swapped {
		t.Errorf("CompareAndSwap(): Expected key %q to be swapped", key)
	}
	v, ok := m.Load(key)
	if !ok || v != swap {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", swap, key, v)
	}
}

func TestCompareAndDelete(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	deleted := m.CompareAndDelete(key, 43)
	if deleted {
		t.Errorf("CompareAndDelete(): Expected key %q to not be deleted", key)
	}
	deleted = m.CompareAndDelete(key, value)
	if !deleted {
		t.Errorf("CompareAndDelete(): Expected key %q to be deleted", key)
	}
	actual, ok := m.Load(key)
	if ok {
		t.Errorf("Load(): Expected key %q to be deleted", key)
	}
	if actual != 0 {
		t.Errorf("Expected value 0, got %d", actual)
	}
}

func TestKeysValues(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Store(i, i)
	}
	sumKeys := 0
	keys := m.Keys()
	for _, key := range keys {
		sumKeys += key
	}
	if len(keys) != 100 {
		t.Errorf("Keys(): Expected 100 keys, got %d", len(keys))
	} else if sumKeys != 4950 {
		t.Errorf("Expected sum 4950, got %d", sumKeys)
	}

	sumValues := 0
	values := m.Values()
	for _, value := range values {
		sumValues += value
	}
	if len(values) != 100 {
		t.Errorf("Values(): Expected 100 values, got %d", len(values))
	}
	if sumValues != 4950 {
		t.Errorf("Expected sum 4950, got %d", sumValues)
	}
}

func TestEntries(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Store(i, i)
	}
	sumKeys := 0
	sumValues := 0
	keys, values := m.Entries()
	for i := 0; i < 100; i++ {
		sumKeys += keys[i]
		sumValues += values[i]
	}
	switch {
	case len(keys) != 100:
		t.Errorf("Keys(): Expected 100 keys, got %d", len(keys))
	case sumKeys != 4950:
		t.Errorf("Expected sum 4950, got %d", sumKeys)
	case len(values) != 100:
		t.Errorf("Values(): Expected 100 values, got %d", len(values))
	case sumValues != 4950:
		t.Errorf("Expected sum 4950, got %d", sumValues)
	}
}

func TestRange(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Store(i, 1)
	}

	var sum int
	m.Range(func(key, value int) bool {
		sum += value
		return true
	})
	if sum != 100 {
		t.Errorf("Range(): Expected sum 4950, got %d", sum)
	}
}

func TestLen(t *testing.T) {
	m := typedmap.New[string, int]()
	if m.Len() != 0 {
		t.Errorf("Len(): Expected length 0, got %d", m.Len())
	}

	m.Store("key1", 42)
	m.Store("key2", 42)
	if m.Len() != 2 {
		t.Errorf("Len(): Expected length 2, got %d", m.Len())
	}
}

func TestUpdate(t *testing.T) {
	m := typedmap.New[string, int]()
	var wg sync.WaitGroup
	n := 100
	loops := 10
	incrementKey := func() {
		for i := 0; i < loops; i++ {
			m.Update("key", func(value int, ok bool) int {
				if !ok {
					panic("Update(): Expected key to be present")
				}
				return value + 1
			})
		}
	}

	m.Store("key", 0)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			incrementKey()
			wg.Done()
		}()
	}

	wg.Wait()

	expectedValue := n * 10
	if value, _ := m.Load("key"); value != expectedValue {
		t.Errorf("Load(): Expected final value to be %d, got %d", expectedValue, value)
	}
}

func TestUpdateRange(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Store(i, 1)
	}

	var sum int
	wg := sync.WaitGroup{}
	wg.Add(1)
	sig := make(chan any)
	go func() {
		<-sig
		defer wg.Done()
		// update range will be called 100 times
		// each time it will add 1 to the value
		// so the final sum should be 100 + 100
		m.UpdateRange(func(k, i int) (int, bool) {
			sum += i
			return i + 1, true
		})
	}()
	sig <- nil
	wg.Wait()
	if sum != 100 {
		t.Errorf("UpdateRange(): Expected sum 100, got %d", sum)
	}
	sum = 0
	// here all values are 2, so the sum should be 200
	m.Range(func(key, value int) bool {
		sum += value
		return true
	})
	if sum != 200 {
		t.Errorf("Range(): Expected sum 200, got %d", sum)
	}
}

func TestTypedMapConcurrentAccessSet(t *testing.T) {
	m := typedmap.New[int, int]()

	// Number of goroutines to spawn
	numGoroutines := 100
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	// starts numGoroutines go routines, the i values is used for content
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			<-ctx.Done()
			// starts numKeys go routines, the j values is used as key
			// so we will have 1000*100 Get, Set and Delete operations
			for j := 0; j < numGoroutines; j++ {
				// uses context done to have all goroutines start at the same time
				_, ok := m.Load(j)
				if !ok {
					m.Store(j, i*i)
				}
				m.Delete(j)
			}
		}(i)
	}
	cancel()
	wg.Wait()
	if m.Len() != 0 {
		t.Errorf("Len(): Expected length 0, got %d", m.Len())
	}
}

func TestTypedMapConcurrentAccessUpdate(t *testing.T) {
	m := typedmap.New[int, int]()

	// Number of goroutines to spawn
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		m.Store(i, 0)
	}
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	// starts numGoroutines go routines, the i values is used for content
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			// uses context done to have all goroutines start at the same time
			<-ctx.Done()
			for j := 0; j < numGoroutines; j++ {
				for j := 0; j < numGoroutines; j++ {
					m.Update(i, func(v int, ok bool) int {
						if !ok {
							panic("Expected key to be present")
						}
						return v + 1
					})
				}
			}
		}(i)
	}
	cancel()
	wg.Wait()
	if m.Len() != numGoroutines {
		t.Errorf("Len(): Expected length 100, got %d", m.Len())
	}
	m.Range(func(k, v int) bool {
		if v != numGoroutines*numGoroutines {
			panic(fmt.Errorf("Expected value %d, got %d", numGoroutines*numGoroutines, v))
		}
		return true
	})

}

func TestTypedMapConcurrentAccessUpdateRange(t *testing.T) {
	m := typedmap.New[int, int]()
	// Number of goroutines to spawn
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		m.Store(i, 0)
	}

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	// starts numGoroutines go routines, the i values is used for content
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			// uses context done to have all goroutines start at the same time
			<-ctx.Done()
			defer wg.Done()
			for j := 0; j < numGoroutines; j++ {
				m.UpdateRange(func(k, v int) (int, bool) {
					return v + 1, true
				})
			}
		}(i)
	}
	cancel()
	wg.Wait()
	if m.Len() != numGoroutines {
		t.Errorf("Len(): Expected length 100, got %d", m.Len())
	}
	m.Range(func(k, v int) bool {
		if v != numGoroutines*numGoroutines {
			panic(fmt.Errorf("Expected value %d, got %d", numGoroutines*numGoroutines, v))
		}
		return true
	})
}

func TestTypedMapConcurrentAccessRange(t *testing.T) {
	m := typedmap.New[int, int]()

	// Number of goroutines to spawn
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		m.Store(i, 0)
	}
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	var count atomic.Int64
	// starts numGoroutines go routines, the i values is used for content
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			defer wg.Done()
			// uses context done to have all goroutines start at the same time
			<-ctx.Done()
			// starts numKeys go routines, the j values is used as key
			// so we will have 1000*100 Get, Set and Delete operations
			for j := 0; j < numGoroutines; j++ {
				m.Range(func(k, v int) bool {
					count.Add(1)
					return true
				})
			}
		}(i)
	}
	cancel()
	wg.Wait()
	c := int(count.Load())
	expect := numGoroutines * numGoroutines * numGoroutines
	if c != expect {
		t.Errorf("Expected count %d, got %d", expect, c)
	}
}

func TestTypedMapConcurrentAccessExclusive(t *testing.T) {
	m := typedmap.New[int, int]()
	// Number of goroutines to spawn
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		m.Store(i, 0)
	}

	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	// starts numGoroutines go routines, the i values is used for content
	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			// uses context done to have all goroutines start at the same time
			<-ctx.Done()
			defer wg.Done()
			for j := 0; j < numGoroutines; j++ {
				m.Exclusive(func(m map[int]int) {
					for k, v := range m {
						m[k] = v + 1
					}
				})
			}
		}(i)
	}
	cancel()
	wg.Wait()
	if m.Len() != numGoroutines {
		t.Errorf("Len(): Expected length 100, got %d", m.Len())
	}
	expect := numGoroutines * numGoroutines
	m.Range(func(k, v int) bool {
		if v != expect {
			panic(fmt.Errorf("Expected value %d, got %d", expect, v))
		}
		return true
	})
}

func TestNotComparableType(t *testing.T) {
	m := typedmap.New[int, []int]()
	m.Store(1, []int{1, 2, 3})
	if m.CompareAndDelete(1, []int{1, 2, 3}) {
		t.Errorf("Expected not comparable type")
	}
	if m.CompareAndSwap(1, []int{1, 2, 3}, []int{1, 2, 3, 4}) {
		t.Errorf("Expected not comparable type")
	}

}
func TestComparableType(t *testing.T) {
	m := typedmap.New[int, int]()
	m.Store(1, 1)
	if !m.CompareAndDelete(1, 1) {
		t.Errorf("Expected to return true")
	}
	m.Store(1, 1)
	if !m.CompareAndSwap(1, 1, 2) {
		t.Errorf("Expected to return true")
	}

}

func TestClear(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Store(i, i)
	}
	if m.Len() != 100 {
		t.Errorf("Len(): Expected length 100, got %d", m.Len())
	}
	m.Clear()
	if m.Len() != 0 {
		t.Errorf("Len(): Expected length 0, got %d", m.Len())
	}
}
