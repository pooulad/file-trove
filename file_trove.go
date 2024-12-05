package main

import (
	"context"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

var UseColor = true

type Storer interface {
	Backup(context.Context) error
}

type FileWorker struct {
	Config Config
}

func NewFileWorker(config Config) Storer {
	return &FileWorker{
		Config: config,
	}
}

func main() {
	config, err := ParseConfig()
	if err != nil {
		log.Fatalf(Colorize(ColorRed, "Error parsing configuration: %v"), err)
	}

	if config.Detach {
		if os.Getenv("IS_CHILD") != "true" {
			cmd := exec.Command(os.Args[0], os.Args[1:]...)
			cmd.Env = append(os.Environ(), "IS_CHILD=true")

			nullDevice := "NUL"
			if runtime.GOOS != "windows" {
				nullDevice = "/dev/null"
			}
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin, _ = os.Open(nullDevice)

			if err := cmd.Start(); err != nil {
				log.Fatalf(Colorize(ColorRed, "Failed to start detached process: %v"), err)
			}

			os.Exit(0)
		}

		UseColor = false

		logFile, err := os.OpenFile(config.LogFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf(Colorize(ColorRed, "Failed to open log file: %v"), err)
		}
		defer logFile.Close()

		log.SetOutput(logFile)
		log.Println("Logging to file:", config.LogFilePath)
	} else {
		log.Println("Logging to terminal")
	}

	storer := NewFileWorker(*config)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup

	switch config.Operation {
	case OperationBackup:
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := storer.Backup(ctx); err != nil {
				log.Printf(Colorize(ColorRed, "Backup failed: %v"), err)
			} else {
				log.Println(Colorize(ColorGreen, "Backup completed successfully"))
			}
		}()
		// case OperationRestore:
		// 	if err := storer.Restore(); err != nil {
		// 		log.Fatalf(Colorize(ColorRed, "Restore failed: %v", err))
		// 	}
	}

	wg.Wait()
}
