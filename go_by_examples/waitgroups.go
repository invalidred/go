package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Workder %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Workder %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(j int) {
			defer wg.Done()
			worker(j)
		}(i)
	}

	wg.Wait()
}
