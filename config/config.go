package config

import (
	"forensic-platform/utils"
)

type Config struct {
	LogLevel string
	MaxDepth int // Maximum directory recursion depth
}

func Load(logLevel string) (*Config, error) {
	cfg := &Config{
		LogLevel: logLevel,
		MaxDepth: 5, // Default max depth for scanning
	}

	// Set up logger based on config
	utils.SetupLogger(cfg.LogLevel)

	return cfg, nil
}
