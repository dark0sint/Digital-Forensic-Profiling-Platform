package analyzer

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"forensic-platform/config"
	"forensic-platform/utils"
)

func ScanDirectory(dir string, cfg *config.Config) (*ProfileReport, error) {
	utils.Logger.Info("Starting scan of directory: %s", dir)

	var files []FileInfo
	totalSize := int64(0)
	fileTypes := make(map[string]int)
	timeline := make(map[time.Time]int) // Aggregated by day

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			utils.Logger.Warn("Error accessing %s: %v", path, err)
			return nil // Continue scanning
		}

		if !info.IsDir() {
			fileType := getFileType(info.Name())
			files = append(files, FileInfo{
				Name:    info.Name(),
				Path:    path,
				Size:    info.Size(),
				ModTime: info.ModTime(),
				Owner:   "unknown", // In real forensics, query os/user
				Type:    fileType,
			})

			totalSize += info.Size()
			fileTypes[fileType]++
			day := time.Date(info.ModTime().Year(), info.ModTime().Month(), info.ModTime().Day(), 0, 0, 0, 0, info.ModTime().Location())
			timeline[day]++
		}

		// Limit depth (simplified check)
		depth := strings.Count(path, string(os.PathSeparator)) - strings.Count(dir, string(os.PathSeparator))
		if depth > cfg.MaxDepth {
			return filepath.SkipDir
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	// Build timeline entries (sorted by date)
	var timelineEntries []TimelineEntry
	for date, count := range timeline {
		timelineEntries = append(timelineEntries, TimelineEntry{Date: date, FileCount: count})
	}
	sort.Slice(timelineEntries, func(i, j int) bool {
		return timelineEntries[i].Date.Before(timelineEntries[j].Date)
	})

	report := &ProfileReport{
		TargetDir:  dir,
		TotalFiles: len(files),
		TotalSize:  totalSize,
		FileTypes:  fileTypes,
		Timeline:   timelineEntries,
		ScanTime:   time.Now(),
	}

	utils.Logger.Info("Scan complete: %d files analyzed", len(files))
	return report, nil
}

func getFileType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".txt", ".md":
		return "text"
	case ".jpg", ".png", ".gif":
		return "image"
	case ".pdf":
		return "document"
	case ".exe", ".bin":
		return "executable"
	default:
		return "other"
	}
}
