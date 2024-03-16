package typedmap_test

import (
	"sync"
	"testing"

	"github.com/thetechpanda/typedmap"
)

func TestNew(t *testing.T) {
	m := typedmap.New[string, int]()
	if m.Len() != 0 {
		t.Errorf("Expected empty map, got map with length %d", m.Len())
	}
}

func TestSetAndGet(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42

	m.Set(key, value)
	v, ok := m.Get(key)

	if !ok || v != value {
		t.Errorf("Expected value %d for key %q, got value %d", value, key, v)
	}
}

func TestDelete(t *testing.T) {
	m := typedmap.New[string, int]()
	key := "key"
	m.Set(key, 42)
	m.Delete(key)

	if _, ok := m.Get(key); ok {
		t.Errorf("Expected key %q to be deleted", key)
	}
}

func TestLen(t *testing.T) {
	m := typedmap.New[string, int]()
	if m.Len() != 0 {
		t.Errorf("Expected length 0, got %d", m.Len())
	}

	m.Set("key1", 42)
	m.Set("key2", 42)
	if m.Len() != 2 {
		t.Errorf("Expected length 2, got %d", m.Len())
	}
}

func TestConcurrentAccess(t *testing.T) {
	m := typedmap.New[int, int]()
	var wg sync.WaitGroup
	n := 1000
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(i, 0)
		}(i)
	}

	wg.Wait()

	if m.Len() != n {
		t.Errorf("Expected length %d, got %d", n, m.Len())
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
					t.Fatalf("Expected key to be present")
				}
				return value + 1
			})
		}
	}

	m.Set("key", 0)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			incrementKey()
			wg.Done()
		}()
	}

	wg.Wait()

	expectedValue := n * 10
	if value, _ := m.Get("key"); value != expectedValue {
		t.Errorf("Expected final value to be %d, got %d", expectedValue, value)
	}
}

func TestUpdateRange(t *testing.T) {
	m := typedmap.New[int, int]()
	for i := 0; i < 100; i++ {
		m.Set(i, 1)
	}

	var sum int
	fwg := sync.WaitGroup{}
	fwg.Add(1)
	first := make(chan bool)
	go func() {
		<-first
		defer fwg.Done()
		m.UpdateRange(func(k, i int) (int, bool) {
			sum += i
			return i + 1, true
		})
	}()
	first <- true
	fwg.Wait()
	if sum != 100 {
		t.Errorf("Expected sum 100, got %d", sum)
	}
	sum = 0
	m.Range(func(key, value int) bool {
		sum += value
		return true
	})
	if sum != 200 {
		t.Errorf("Expected sum 200, got %d", sum)
	}
}
