package main
import(
	"fmt"
	"strings"
	"sync"
)

func countWord(word string, total *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	if word != "" {
		mutex.Lock()
		*total = *total + 1
		mutex.Unlock()
	}
}

func main() {
	text := "The quick brown fox jumps over the lazy dog"
	words := strings.Split(text, " ")
	total := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, word := range words {
		wg.Add(1)
		go countWord(
			word,
			&total,
			&mutex,
			&wg,
		)
	}
	wg.Wait()
	fmt.Println("Number of words:", total)
}