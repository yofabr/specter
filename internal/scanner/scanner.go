package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Config struct {
	Target    string
	StartPort int
	EndPort   int
	Workers   int
}

type Result struct {
	Port  int
	State string
}

type Scanner struct {
	config Config
}

func NewScanner(cfg Config) *Scanner {
	return &Scanner{config: cfg}
}

func (s *Scanner) Scan() []Result {
	var wg sync.WaitGroup
	ports := make(chan int, s.config.EndPort-s.config.StartPort+1)
	results := make(chan Result, s.config.EndPort-s.config.StartPort+1)

	for i := s.config.StartPort; i <= s.config.EndPort; i++ {
		ports <- i
	}
	close(ports)

	for i := 0; i < s.config.Workers; i++ {
		wg.Add(1)
		go s.worker(ports, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	var openPorts []Result
	for r := range results {
		if r.State == "open" {
			openPorts = append(openPorts, r)
		}
	}

	return openPorts
}

func (s *Scanner) ScanWithCallback(callback func(Result)) {
	var wg sync.WaitGroup
	ports := make(chan int, s.config.EndPort-s.config.StartPort+1)
	results := make(chan Result, s.config.EndPort-s.config.StartPort+1)

	for i := s.config.StartPort; i <= s.config.EndPort; i++ {
		ports <- i
	}
	close(ports)

	for i := 0; i < s.config.Workers; i++ {
		wg.Add(1)
		go s.worker(ports, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		if r.State == "open" {
			callback(r)
		}
	}
}

func (s *Scanner) worker(ports <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		address := s.config.Target + ":" + strconv.Itoa(p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			results <- Result{Port: p, State: "open"}
			conn.Close()
		} else {
			results <- Result{Port: p, State: "closed"}
		}
	}
}

func PrintResult(r Result) {
	fmt.Printf("Port %d: %s\n", r.Port, r.State)
}
