package main

import (
	"context"
	"fmt"
	"time"
)

// SleepAfter blocks the goroutine using time.After
// + Simple and easy to use
// + Non-blocking for CPU (efficient)
// - Creates a new timer every call (minor overhead)
func SleepAfter(d time.Duration) {
	<-time.After(d)
}

// SleepTicker blocks the goroutine using a one-time tick
// + Works for a single event
// - Uses a ticker which is designed for repeated events (less efficient)
// - Slight overhead due to ticker creation
func SleepTicker(d time.Duration) {
	ticker := time.NewTicker(d)
	<-ticker.C
	ticker.Stop()
}

// SleepInterruptible allows sleeping that can be interrupted via context
// + Can be canceled, very flexible
// + Efficient, does not block CPU
// - Slightly more complex than basic sleep
func SleepInterruptible(ctx context.Context, d time.Duration) error {
	select {
	case <-time.After(d):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SleepProgress shows sleep progress at a given interval
// + Good for progress display or feedback
// + Non-blocking if using time.Sleep for interval
// - Slight overhead due to repeated printing and sleeping
func SleepProgress(d time.Duration, interval time.Duration) {
	start := time.Now()
	for {
		elapsed := time.Since(start)
		if elapsed >= d {
			break
		}
		fmt.Printf("Elapsed: %v / %v\n", elapsed.Truncate(time.Millisecond), d)
		time.Sleep(interval)
	}
}

// BusySleep blocks execution using a busy loop (CPU-intensive)
// + Demonstrates fully blocking without timers
// - Very CPU heavy, not recommended in production
// - No interruption possible
func BusySleep(d time.Duration) {
	end := time.Now().Add(d)
	for time.Now().Before(end) {
		// busy-wait loop
	}
}

func main() {
	fmt.Println("SleepAfter: Asleep...")
	SleepAfter(2 * time.Second)
	fmt.Println("Wake up!")

	fmt.Println("\nSleepTicker: Asleep...")
	SleepTicker(2 * time.Second)
	fmt.Println("Wake up!")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	fmt.Println("\nSleepInterruptible: Trying to sleep 5s with 2s timeout...")
	err := SleepInterruptible(ctx, 5*time.Second)
	if err != nil {
		fmt.Println("Interrupted:", err)
	} else {
		fmt.Println("Slept successfully")
	}

	fmt.Println("\nSleepProgress: 3s sleep with 1s intervals")
	SleepProgress(3*time.Second, 1*time.Second)
	fmt.Println("Done!")

	fmt.Println("\nBusySleep: 2s sleep (CPU busy)")
	BusySleep(2 * time.Second)
	fmt.Println("Done!")
}
