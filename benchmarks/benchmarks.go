package benchmarks

import (
	"context"
	"sync"
	"testing"
)

// use is an helper function that does nothing with the input.
func noop(x ...any) {
	_ = x
}

// benchmarkConcurrentInt is a helper function to benchmark concurrent operations on an integer.
func benchmarkConcurrentInt(b *testing.B, f func(n, i, j int)) {
	numGoroutines := 100
	numOperations := 100
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
