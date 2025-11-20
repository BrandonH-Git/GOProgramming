package main

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func resolver(Website string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	ips, err := net.LookupIP(Website)
	if err != nil {
		log.Println(err)
		return
	}
	for _, ip := range ips {
		if ip.To4() != nil {
			results <- fmt.Sprintf("%s -> %s", Website, ip.String())
		}

	}
}
func main() {
	Websites := []string{"www.example.com", "www.iana.org", "www.golang.org"}
	var wg sync.WaitGroup
	results := make(chan string, len(Websites))
	for i := range Websites {
		wg.Add(1)
		go resolver(Websites[i], &wg, results)
	}
	go func() {
		for res := range results {
			fmt.Println(res)
		}
	}()

	wg.Wait()
	close(results)
}
