package main

import (
	"flag"
	"fmt"

	"specter/internal/scanner"
)

func main() {
	target := flag.String("target", "127.0.0.1", "Target IP address")
	startPort := flag.Int("start", 1, "Start port")
	endPort := flag.Int("end", 1024, "End port")
	workers := flag.Int("workers", 100, "Number of workers")
	flag.Parse()

	cfg := scanner.Config{
		Target:    *target,
		StartPort: *startPort,
		EndPort:   *endPort,
		Workers:   *workers,
	}

	s := scanner.NewScanner(cfg)

	fmt.Printf("Scanning %s (ports %d-%d)...\n", *target, *startPort, *endPort)
	s.ScanWithCallback(scanner.PrintResult)
	fmt.Println("Scan complete.")
}
