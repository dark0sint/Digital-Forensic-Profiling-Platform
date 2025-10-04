package analyzer

import (
	"time"
)

type FileInfo struct {
	Name     string    `json:"name"`
	Path     string    `json:"path"`
	Size     int64     `json:"size_bytes"`
	ModTime  time.Time `json:"modification_time"`
	Owner    string    `json:"owner"` // Simulated; in real use, use os.User
	Type     string    `json:"type"`  // e.g., "text", "image", based on extension
}

type ProfileReport struct {
	TargetDir    string         `json:"target_directory"`
	TotalFiles   int            `json:"total_files"`
	TotalSize    int64          `json:"total_size_bytes"`
	FileTypes    map[string]int `json:"file_type_distribution"`
	Timeline     []TimelineEntry `json:"modification_timeline"`
	ScanTime     time.Time      `json:"scan_timestamp"`
}

type TimelineEntry struct {
	Date     time.Time `json:"date"`
	FileCount int      `json:"file_count"`
}
