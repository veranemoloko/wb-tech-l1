package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Counter is a generic interface for any counter implementation.
type Counter interface {
	Inc()
	Value() uint64
}

// -----Mutex-----
type CounterWithMutex struct {
	val uint64
	mu  sync.Mutex
}

func (c *CounterWithMutex) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *CounterWithMutex) Value() uint64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}

// -----Atomic-----
type CounterWithAtomic struct {
	val uint64
}

func (c *CounterWithAtomic) Inc() {
	atomic.AddUint64(&c.val, 1)
}

func (c *CounterWithAtomic) Value() uint64 {
	return atomic.LoadUint64(&c.val)
}

// benchmarkCounter launches multiple goroutines, each performing increments
// on the given counter. It waits for all goroutines to finish and returns
// the final counter value and execution time.
func benchmarkCounter(counter Counter, goroutines, iterations int) (uint64, time.Duration) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				counter.Inc()
			}
		}()
	}

	wg.Wait()
	return counter.Value(), time.Since(start)
}

func main() {
	workersCnt := 100
	iterations := 100000

	mc := &CounterWithMutex{}
	val1, dur1 := benchmarkCounter(mc, workersCnt, iterations)
	fmt.Printf("Mutex:   value=%d, time=%s\n", val1, dur1)

	ac := &CounterWithAtomic{}
	val2, dur2 := benchmarkCounter(ac, workersCnt, iterations)
	fmt.Printf("Atomic:  value=%d, time=%s\n", val2, dur2)

	// Atomic is usually faster than Mutex,
	// because it does not require locking/unlocking,
	// only a single CPU atomic instruction.
}
