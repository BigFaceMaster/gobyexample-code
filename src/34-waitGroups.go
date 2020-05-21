package main

import (
	"fmt"
	"sync"
	"time"
)

func waitWorker(id int, wg *sync.WaitGroup)  {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Workder %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i<= 5; i++ {
		wg.Add(1)
		go waitWorker(i, &wg)
	}

	wg.Wait()
}
