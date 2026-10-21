package main
import(
	"fmt"
	"sync"
)

func generate(numbers chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(numbers)
	for i := 1; i <= 10; i++ {
		numbers <- i
	}
}

func double(numbers chan int, doubled chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(doubled)
	for number := range numbers {
		doubled <- number * 2
	}
}

func printNumbers(doubled chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range doubled {
		fmt.Println(number)
	}
}

func main(){
	numbers := make(chan int, 10)
	doubled := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(3)
	go generate(numbers, &wg)
	go double(numbers, doubled, &wg)
	go printNumbers(doubled, &wg)
	wg.Wait()
}