package main

import (
	"fmt"
	"sync"
)

func hello(name string, wg *sync.WaitGroup) {
	fmt.Println("Hello " + name + "!")
	defer wg.Done()
}

func main() {
	var wg sync.WaitGroup
	names := []string{"Alice", "Bob", "Charlie", "Dave", "John"}
	for i := range names {
		wg.Add(1)
		go hello(names[i], &wg)
	}
	wg.Wait()
}
