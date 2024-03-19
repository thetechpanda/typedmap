package typedmap_test

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/thetechpanda/typedmap"
)

func ExampleTypedMap() {
	m := typedmap.New[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	// v is int
	fmt.Println("v:", v, "ok:", ok)
}

func ExampleSyncMap() {
	m := typedmap.NewSyncMap[string, int]()
	key := "key"
	value := 42
	m.Store(key, value)
	v, ok := m.Load(key)
	// v is int
	fmt.Println("v:", v, "ok:", ok)
}

func Example() {
	m := typedmap.New[string, int]() // or typedmap.NewSyncMap[string, int]()
	for i := 0; i < 10; i++ {
		m.Store(fmt.Sprintf("key%d", i), i)
	}

	context, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}

	wg.Add(1)
	var totalReads = 0
	read := func() {
		i := 0
		for {
			select {
			case <-context.Done():
				wg.Done()
				fmt.Println("read() exit")
				return
			default:
				m.Load(fmt.Sprintf("key%d", i))
				// do something with the value
				i = (i + 1) % 10
				totalReads++
			}
		}
	}

	wg.Add(1)
	var totalWrites = 0
	write := func() {
		i := 0
		for {
			select {
			case <-context.Done():
				wg.Done()
				fmt.Println("write(): exit")
				return
			default:
				m.Swap(fmt.Sprintf("key%d", i), i)
				// do something with the value
				i = (i + 1) % 10
				totalWrites++
			}
		}
	}

	go read()
	go write()

	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	fmt.Printf("read count: %d write count: %d\n", totalReads, totalWrites)
}
