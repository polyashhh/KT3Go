package main

import (
	"fmt"
	"sync"
)

func sumPart(numbers []int, result chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	result <- sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make(chan int, 2)
	var wg sync.WaitGroup
	wg.Add(2)
	go sumPart(numbers[:5], result, &wg)
	go sumPart(numbers[5:], result, &wg)
	wg.Wait()
	firstSum := <-result
	secondSum := <-result
	totalSum := firstSum + secondSum
	fmt.Println("Sum:", totalSum)
}
