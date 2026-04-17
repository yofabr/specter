package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"specter/internal/scanner"

	"github.com/spf13/cobra"
)

var target string
var startPort int
var endPort int
var workers int
var timeout int

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan target for open ports",
	Run: func(cmd *cobra.Command, args []string) {
		cfg := scanner.Config{
			Target:    target,
			StartPort: startPort,
			EndPort:   endPort,
			Workers:   workers,
		}

		if timeout > 0 {
			cfg.Timeout = scanner.ParseTimeout(timeout)
		}

		if err := cfg.Validate(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		s := scanner.NewScanner(cfg)

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		defer signal.Stop(sigChan)

		fmt.Printf("Scanning %s (ports %d-%d)...\n", target, startPort, endPort)

		go func() {
			<-sigChan
			fmt.Println("\nReceived interrupt, shutting down...")
			os.Exit(130)
		}()

		s.ScanWithCallback(func(r scanner.Result) {
			fmt.Printf("Port %d: %s\n", r.Port, r.State)
		})

		fmt.Println("Scan complete.")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&target, "target", "t", "127.0.0.1", "Target IP address")
	scanCmd.Flags().IntVarP(&startPort, "start", "s", 1, "Start port")
	scanCmd.Flags().IntVarP(&endPort, "end", "e", 1024, "End port")
	scanCmd.Flags().IntVarP(&workers, "workers", "w", 100, "Number of workers")
	scanCmd.Flags().IntVarP(&timeout, "timeout", "", 2, "Connection timeout in seconds")
}
