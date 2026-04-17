package scanner

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

type Config struct {
	Target    string
	StartPort int
	EndPort   int
	Workers   int
	Timeout   time.Duration
}

func (c Config) Validate() error {
	if c.Target == "" {
		return fmt.Errorf("target is required")
	}
	if ip := net.ParseIP(c.Target); ip == nil {
		return fmt.Errorf("invalid target IP address: %s", c.Target)
	}
	if c.StartPort < 1 || c.StartPort > 65535 {
		return fmt.Errorf("start port must be between 1 and 65535, got %d", c.StartPort)
	}
	if c.EndPort < 1 || c.EndPort > 65535 {
		return fmt.Errorf("end port must be between 1 and 65535, got %d", c.EndPort)
	}
	if c.StartPort > c.EndPort {
		return fmt.Errorf("start port (%d) cannot be greater than end port (%d)", c.StartPort, c.EndPort)
	}
	if c.Workers < 1 {
		return fmt.Errorf("workers must be at least 1, got %d", c.Workers)
	}
	if c.Workers > 10000 {
		return fmt.Errorf("workers cannot exceed 10000, got %d", c.Workers)
	}
	if c.Timeout < 0 {
		return fmt.Errorf("timeout cannot be negative")
	}
	return nil
}

type Result struct {
	Port  int
	State string
}

type Scanner struct {
	config Config
}

func NewScanner(cfg Config) *Scanner {
	if cfg.Timeout == 0 {
		cfg.Timeout = 2 * time.Second
	}
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

	wg.Wait()
	close(results)

	var openPorts []Result
	for r := range results {
		openPorts = append(openPorts, r)
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

	wg.Wait()
	close(results)

	for r := range results {
		callback(r)
	}
}

func (s *Scanner) worker(ports <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for p := range ports {
		address := s.config.Target + ":" + strconv.Itoa(p)
		conn, err := net.DialTimeout("tcp", address, s.config.Timeout)
		if err == nil {
			results <- Result{Port: p, State: "open"}
			conn.Close()
		}
	}
}

func PrintResult(r Result) {
	fmt.Printf("Port %d: %s\n", r.Port, r.State)
}

func ParseTimeout(seconds int) time.Duration {
	return time.Duration(seconds) * time.Second
}
