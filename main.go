package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"forensic-platform/analyzer"
	"forensic-platform/config"
	"forensic-platform/utils"
)

func main() {
	// Parse CLI flags
	dir := flag.String("dir", "", "Target directory to scan")
	output := flag.String("output", "report.json", "Output file for the profile report")
	logLevel := flag.String("log", "info", "Log level (debug, info, error)")
	flag.Parse()

	if *dir == "" {
		fmt.Println("Error: --dir flag is required.")
		os.Exit(1)
	}

	// Load configuration
	cfg, err := config.Load(*logLevel)
	if err != nil {
		utils.Logger.Error("Failed to load config: %v", err)
		os.Exit(1)
	}

	// Run forensic analysis
	profile, err := analyzer.ScanDirectory(*dir, cfg)
	if err != nil {
		utils.Logger.Error("Analysis failed: %v", err)
		os.Exit(1)
	}

	// Output report as JSON
	reportBytes, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		utils.Logger.Error("Failed to marshal report: %v", err)
		os.Exit(1)
	}

	err = os.WriteFile(*output, reportBytes, 0644)
	if err != nil {
		utils.Logger.Error("Failed to write report: %v", err)
		os.Exit(1)
	}

	utils.Logger.Info("Profile report generated: %s", *output)
	fmt.Printf("Scan complete. Report saved to %s\n", *output)
}
