package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func SetupLogger(level string) {
	Logger = log.New(os.Stdout, "[ForensicPlatform] ", log.LstdFlags|log.Lshortfile)
	switch level {
	case "debug":
		Logger.SetFlags(log.LstdFlags | log.Lshortfile)
	case "error":
		// Only log errors (simplified; in real code, filter)
	default:
		// Info level
	}
}

func (l *log.Logger) Debug(msg string, args ...interface{}) {
	l.Printf("[DEBUG] "+msg, args...)
}

func (l *log.Logger) Info(msg string, args ...interface{}) {
	l.Printf("[INFO] "+msg, args...)
}

func (l *log.Logger) Warn(msg string, args ...interface{}) {
	l.Printf("[WARN] "+msg, args...)
}

func (l *log.Logger) Error(msg string, args ...interface{}) {
	l.Printf("[ERROR] "+msg, args...)
}
