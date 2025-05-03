package cuncurrency

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
func RunConcurrently() {
	fmt.Println("\nRunning Concurrently:")
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("fib(35):", fib(35))
	}()
	go func() {
		defer wg.Done()
		fmt.Println("fib(36):", fib(36))
	}()

	wg.Wait()
	fmt.Println("Time taken:", time.Since(start))
}

func RunSequentially() {
	fmt.Println("Running Sequentially:")
	start := time.Now()
	fmt.Println("fib(35):", fib(35))
	fmt.Println("fib(36):", fib(36))
	fmt.Println("Time taken:", time.Since(start))
}

func RunInParallel() {
	fmt.Println("\nRunning in Parallel (GOMAXPROCS):")
	runtime.GOMAXPROCS(2) // Tell Go to use 2 OS threads (CPU cores)
	var wg sync.WaitGroup
	start := time.Now()

	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("fib(35):", fib(35))
	}()
	go func() {
		defer wg.Done()
		fmt.Println("fib(36):", fib(36))
	}()

	wg.Wait()
	fmt.Println("Time taken:", time.Since(start))
}
