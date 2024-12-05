package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	OperationBackup  = "backup"
	OperationRestore = "restore"
)

type Config struct {
	Operation   string
	Source      string
	Destination string
	Period      int
	Compress    bool
	Detach      bool
	LogFilePath string
	Exclude     []string
}

func ParseConfig() (*Config, error) {
	operation := flag.String("operation", "backup", "Operation type")
	source := flag.String("src", "", "Source of backup operation")
	destination := flag.String("dest", "", "Destination of backup operation")
	period := flag.Int("period", 0, "Period of backup operation each time by minute")
	compress := flag.Bool("compress", false, "Enable compression")
	detach := flag.Bool("detach", false, "Enable detach run in background and log in new file instead of terminal")
	logFilePath := flag.String("log", "file-trove.log", "Log file path. default is .")
	exclude := flag.String("exclude", "", "Comma-separated list of files/directories to exclude")

	flag.Parse()

	if *operation == "" {
		return nil, fmt.Errorf("operation must be specified")
	}

	if *operation != OperationBackup && *operation != OperationRestore {
		return nil, fmt.Errorf("invalid operation: %s. Allowed values are 'backup' or 'restore'", *operation)
	}

	if *source == "" || *destination == "" {
		return nil, fmt.Errorf("source and destination must be provided")
	}

	config := &Config{
		Source:      *source,
		Destination: *destination,
		Compress:    *compress,
		Detach:      *detach,
		LogFilePath: *logFilePath,
		Period:      *period,
		Operation:   *operation,
	}

	if *exclude != "" {
		config.Exclude = strings.Split(*exclude, ",")
	}

	return config, nil
}
