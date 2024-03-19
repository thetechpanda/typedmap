package mutex_test

import (
	"context"
	"fmt"
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap"
	"github.com/thetechpanda/typedmap/internal/mutex"
)

func TestNew(t *testing.T) {
	m := mutex.New[string, int](nil)
	if m.Len() != 0 {
		t.Errorf("Len(): Expected a new map, got map with length %d", m.Len())
	}
}
func TestNewWithMap(t *testing.T) {
	data := map[string]int{
		"key1": 42,
		"key2": 43,
	}
	m := mutex.New(data)
	v, ok := m.Load("key1")
	if !ok || v != 42 {
		t.Errorf("Load(): Expected value 42 for key %q, got value %d", "key1", v)
	}
	if !m.Has("key2") {
		t.Errorf("Has(): Expected key %q to be present", "key2")
	}

	v, ok = m.Load("key2")
	if !ok || v != 43 {
		t.Errorf("Load(): Expected value 43 for key %q, got value %d", "key2", v)
	}
}

func TestLoad(t *testing.T) {
	m := mutex.New[string, int](nil)
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok || v != value {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", value, key, v)
	}
}

func TestStore(t *testing.T) {
	m := mutex.New[string, int](nil)
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok || v != value {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", value, key, v)
	}
}

func TestLoadOrStore(t *testing.T) {
	m := mutex.New[string, int](nil)
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
	m := mutex.New[string, int](nil)
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
	m := mutex.New[string, int](nil)
	key := "key"
	value := 42
	m.Store(key, value)
	m.Delete(key)
	if _, ok := m.Load(key); ok {
		t.Errorf("Load(): Expected key %q to be deleted", key)
	}
}

func TestSwap(t *testing.T) {
	m := mutex.New[string, int](nil)
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
	m := mutex.New[string, int](nil)
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
	m := mutex.New[string, int](nil)
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

	if len(m.Keys()) != 0 {
		t.Errorf("Keys(): Expected empty keys, got %v", m.Keys())
	}

	if len(m.Values()) != 0 {
		t.Errorf("Values(): Expected empty values, got %v", m.Values())
	}

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

	zK, zV := m.Entries()
	if len(zK) != 0 {
		t.Errorf("Entries(): Expected empty keys, got %v", zK)
	}
	if len(zV) != 0 {
		t.Errorf("Entries(): Expected empty values, got %v", zV)
	}

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
	sum = 0
	m.Range(func(key, value int) bool {
		sum++
		return false
	})
	if sum != 1 {
		t.Errorf("Range(): Expected sum 1, got %d", sum)
	}
}

func TestLen(t *testing.T) {
	m := mutex.New[string, int](nil)
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
	m := mutex.New[string, int](nil)
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
			sum++
			return 2, true
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
	sum = 0
	// here all values are 2, so the sum should be 200
	m.Range(func(key, value int) bool {
		sum++
		return false
	})
	if sum != 1 {
		t.Errorf("Range(): Expected sum 1, got %d", sum)
	}

	sum = 0
	var mapKey int
	m.UpdateRange(func(k, i int) (int, bool) {
		sum++
		mapKey = k
		return 0, false
	})

	if sum != 1 {
		t.Errorf("UpdateRange(): Expected sum 1, got %d", sum)
	}

	if v, ok := m.Load(mapKey); !ok {
		t.Errorf("Load(): Expected key %d to be present", mapKey)
	} else if v != 2 {
		t.Errorf("Load(): Expected value 2, got %d", v)
	}
}

func TestConcurrentAccessSet(t *testing.T) {
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

func TestConcurrentAccessUpdate(t *testing.T) {
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

func TestConcurrentAccessUpdateRange(t *testing.T) {
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

func TestConcurrentAccessRange(t *testing.T) {
	m := typedmap.New[int, int]()

	// Number of goroutines to spawn
	numGoroutines := 100
	for i := 0; i < numGoroutines; i++ {
		m.Store(i, 0)
	}
	var wg sync.WaitGroup
	wg.Add(numGoroutines)
	ctx, cancel := context.WithCancel(context.Background())
	var count int
	var mu sync.Mutex
	increment := func() {
		mu.Lock()
		defer mu.Unlock()
		count++
	}
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
					increment()
					return true
				})
			}
		}(i)
	}
	cancel()
	wg.Wait()
	expect := numGoroutines * numGoroutines * numGoroutines
	if count != expect {
		t.Errorf("Expected count %d, got %d", expect, count)
	}
}

func TestConcurrentAccessExclusive(t *testing.T) {
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

	if m.CompareAndSwap(1, 3, 4) {
		t.Errorf("Expected not to return true")
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

func testValues[K comparable, V comparable](t *testing.T, key K, value V) {
	m := typedmap.New[K, V]()
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok {
		t.Errorf("Load(): Expected key %v to be present", key)
	}
	if v != value {
		t.Errorf("Load(): Expected value %v for key %v, got value %v", value, key, v)
	}

	m.Delete(key)

	if _, loaded := m.LoadOrStore(key, value); loaded {
		t.Errorf("LoadOrStore(): Expected key %v not to be stored", key)
	}
	if !m.CompareAndSwap(key, value, value) {
		t.Errorf("CompareAndSwap(): Expected key %v to be swapped", key)
	}
	if !m.CompareAndDelete(key, value) {
		t.Errorf("CompareAndDelete(): Expected key %v to be deleted", key)
	}
	m.CompareAndDelete(key, value)
	m.Swap(key, value)
	m.LoadAndDelete(key)

}

func TestValues(t *testing.T) {

	// Test values of different types
	var b bool = true
	var s string = "string"
	var i int = 1
	var f float64 = 1.1
	var c complex128 = 1 + 1i
	var r rune = 'r'
	var by byte = byte('b')
	var st = struct {
		A int
	}{A: 1}

	// Pointer to values
	var bP *bool = &b
	var sP *string = &s
	var iP *int = &i
	var fP *float64 = &f
	var cP *complex128 = &c
	var rP *rune = &r
	var byP *byte = &by
	var stP *struct {
		A int
	} = &st

	// Nil pointers
	var bN *bool = nil
	var sN *string = nil
	var iN *int = nil
	var fN *float64 = nil
	var cN *complex128 = nil
	var rN *rune = nil
	var byN *byte = nil
	var stN *struct {
		A int
	} = nil

	testValues(t, b, b)
	testValues(t, s, s)
	testValues(t, i, i)
	testValues(t, f, f)
	testValues(t, c, c)
	testValues(t, r, r)
	testValues(t, by, by)
	testValues(t, st, st)

	testValues(t, b, bP)
	testValues(t, s, sP)
	testValues(t, i, iP)
	testValues(t, f, fP)
	testValues(t, c, cP)
	testValues(t, r, rP)
	testValues(t, by, byP)
	testValues(t, st, stP)

	testValues(t, bP, b)
	testValues(t, sP, s)
	testValues(t, iP, i)
	testValues(t, fP, f)
	testValues(t, cP, c)
	testValues(t, rP, r)
	testValues(t, byP, by)
	testValues(t, stP, st)

	testValues(t, bN, b)
	testValues(t, sN, s)
	testValues(t, iN, i)
	testValues(t, fN, f)
	testValues(t, cN, c)
	testValues(t, rN, r)
	testValues(t, byN, by)
	testValues(t, stN, st)

	testValues(t, b, bN)
	testValues(t, s, sN)
	testValues(t, i, iN)
	testValues(t, f, fN)
	testValues(t, c, cN)
	testValues(t, r, rN)
	testValues(t, by, byN)
	testValues(t, st, stN)

	testValues(t, bP, bN)
	testValues(t, sP, sN)
	testValues(t, iP, iN)
	testValues(t, fP, fN)
	testValues(t, cP, cN)
	testValues(t, rP, rN)
	testValues(t, byP, byN)
	testValues(t, stP, stN)

}
