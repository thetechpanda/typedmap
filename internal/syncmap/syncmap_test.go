package syncmap_test

import (
	"context"
	"reflect"
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap/internal/syncmap"
)

func TestSyncMapLoad(t *testing.T) {
	m := syncmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok || v != value {
		t.Errorf("Load(): Expected value %d for key %q, got value %d", value, key, v)
	}
	v, ok = m.Load("not-existent")
	if ok {
		t.Errorf("Load(): Expected no value for key %q, got value %d", key, v)
	}
}

func TestSyncMapLoadOrStore(t *testing.T) {
	m := syncmap.New[string, int]()
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

	mp := syncmap.New[string, *int]()
	mp.Store(key, nil)
	var valueP *int = new(int)
	if actual, _ := mp.LoadOrStore(key, valueP); actual != nil {
		t.Errorf("LoadOrStore(): Expected to store nil value")
	}

}

func TestSyncMapLoadAndDelete(t *testing.T) {
	m := syncmap.New[string, int]()
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

func TestSyncMapDelete(t *testing.T) {
	m := syncmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	m.Delete(key)
	if _, ok := m.Load(key); ok {
		t.Errorf("Load(): Expected key %q to be deleted", key)
	}
}

func TestSyncMapSwap(t *testing.T) {
	m := syncmap.New[string, int]()
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

func TestSyncMapCompareAndSwap(t *testing.T) {
	m := syncmap.New[string, int]()
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

func TestSyncMapCompareAndDelete(t *testing.T) {
	m := syncmap.New[string, int]()
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

func TestSyncMapConcurrentAccessStore(t *testing.T) {
	m := syncmap.New[int, int]()

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
}

func TestSyncMapNotComparableTypeCompareAndDelete(t *testing.T) {
	pass := true
	defer func() {
		recover()
		pass = false
	}()
	m := syncmap.New[int, []int]()
	m.Store(1, []int{1, 2, 3})
	if m.CompareAndDelete(1, []int{1, 2, 3}) {
		t.Errorf("Expected not comparable type")
	}
	if !pass {
		t.Errorf("Expected panic")
	}
}

func TestSyncMapNotComparableTypeCompareAndSwap(t *testing.T) {
	pass := true
	defer func() {
		recover()
		pass = false
	}()
	m := syncmap.New[int, []int]()
	m.Store(1, []int{1, 2, 3})
	if m.CompareAndSwap(1, []int{1, 2, 3}, []int{1, 2, 3, 4}) {
		t.Errorf("Expected not comparable type")
	}
	if pass {
		t.Errorf("Expected panic")
	}
}

func TestSyncMapComparableType(t *testing.T) {
	m := syncmap.New[int, int]()
	m.Store(1, 1)
	if !m.CompareAndDelete(1, 1) {
		t.Errorf("Expected to return true")
	}
	m.Store(1, 1)
	if !m.CompareAndSwap(1, 1, 2) {
		t.Errorf("Expected to return true")
	}

}

func TestSyncMapRange(t *testing.T) {
	m := syncmap.New[int, int]()
	numKeys := 100
	for i := 0; i < numKeys; i++ {
		m.Store(i, i)
	}
	count := 0
	m.Range(func(key int, value int) bool {
		count++
		return true
	})
	if count != numKeys {
		t.Errorf("Expected to iterate over all keys")
	}
	count = 0
	m.Range(func(key int, value int) bool {
		count++
		return false
	})
	if count != 1 {
		t.Errorf("Expected to iterate over all keys")
	}
}

func testValues[K any, V any](t *testing.T, key K, value V) {
	m := syncmap.New[K, V]()
	m.Store(key, value)
	v, ok := m.Load(key)
	if !ok {
		t.Errorf("Load(): Expected key %v to be present", key)
	}

	if !reflect.DeepEqual(v, value) {
		t.Errorf("Load(): Expected value %v for key %v, got value %v", value, key, v)
	}

	m.LoadOrStore(key, value)
	m.CompareAndSwap(key, value, value)
	m.CompareAndDelete(key, value)
	m.LoadAndDelete(key)
	m.Delete(key)
	m.Swap(key, value)
}

func TestValues(t *testing.T) {
	// test pointer to any type with nil value
	var anyV interface{}
	ma := syncmap.New[*interface{}, interface{}]()
	ma.Store(&anyV, anyV)
	ma.LoadOrStore(&anyV, anyV)
	// test pointer to struct with nil value
	var anyS struct{}
	ms := syncmap.New[*struct{}, struct{}]()
	ms.Store(&anyS, anyS)
	ms.LoadOrStore(&anyS, anyS)

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
