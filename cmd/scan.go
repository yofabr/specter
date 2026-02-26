package main

import (
	"fmt"

	"specter/internal/scanner"

	"github.com/spf13/cobra"
)

var target string
var startPort int
var endPort int
var workers int

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

		s := scanner.NewScanner(cfg)

		fmt.Printf("Scanning %s (ports %d-%d)...\n", target, startPort, endPort)
		s.ScanWithCallback(scanner.PrintResult)
		fmt.Println("Scan complete.")
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&target, "target", "t", "127.0.0.1", "Target IP address")
	scanCmd.Flags().IntVarP(&startPort, "start", "s", 1, "Start port")
	scanCmd.Flags().IntVarP(&endPort, "end", "e", 1024, "End port")
	scanCmd.Flags().IntVarP(&workers, "workers", "w", 100, "Number of workers")
}
