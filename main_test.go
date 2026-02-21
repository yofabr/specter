package main

import (
	"flag"
	"sync"
	"testing"
)

func TestWorker(t *testing.T) {
	target := flag.String("target", "127.0.0.1", "Target IP address")
	flag.Parse()

	ports := make(chan int, 3)
	ports <- 80
	ports <- 443
	ports <- 8080
	close(ports)

	var wg sync.WaitGroup
	wg.Add(1)
	go worker(*target, ports, &wg)
	wg.Wait()
}
