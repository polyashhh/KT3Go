package main

import (
	"fmt"
	"sync"
)

type Result struct {
	number    int
	factorial uint64
}

func factorial(n int) uint64 {
	result := uint64(1)
	for i := 1; i <= n; i++ {
		result *= uint64(i)
	}
	return result
}

func worker(tasks chan int, results chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range tasks {
		results <- Result{
			number:    number,
			factorial: factorial(number),
		}
	}
}

func main() {
	tasks := make(chan int, 20)
	results := make(chan Result, 20)
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go worker(tasks, results, &wg)
	}
	for i := 1; i <= 20; i++ {
		tasks <- i
	}
	close(tasks)
	go func() {
		wg.Wait()
		close(results)
	}()
	for result := range results {
		fmt.Println(
			"Factorial of",
			result.number,
			"is",
			result.factorial,
		)
	}
}
