package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	target := flag.String("target", "127.0.0.1", "Target IP address")
	flag.Parse()

	var wg sync.WaitGroup
	ports := make(chan int, 1024)

	for i := 1; i <= 1024; i++ {
		ports <- i
	}
	close(ports)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go worker(*target, ports, &wg)
	}

	wg.Wait()
}

func worker(target string, ports <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		address := target + ":" + strconv.Itoa(p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			fmt.Printf("Port %d: Open\n", p)
			conn.Close()
		}
	}
}
