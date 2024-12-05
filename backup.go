package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	cp "github.com/otiai10/copy"
)

func (f *FileWorker) Backup(ctx context.Context) error {
	sourceBase := filepath.Base(f.Config.Source)
	destPath := filepath.Join(f.Config.Destination, sourceBase)

	backupFunc := func() error {
		if err := cp.Copy(f.Config.Source, destPath); err != nil {
			return fmt.Errorf("failed to copy files: %v", err)
		}

		if f.Config.Compress {
			zipBase := filepath.Base(destPath) + ".zip"
			zipPath := filepath.Join(filepath.Dir(destPath), zipBase)

			log.Printf(Colorize(ColorYellow, "Compressing %s to %s"), destPath, zipPath)
			if err := CompressToZip(destPath, zipPath); err != nil {
				return fmt.Errorf("zip compression failed: %v", err)
			}

			if err := os.RemoveAll(destPath); err != nil {
				log.Printf(Colorize(ColorRed, "Failed to delete original files after compression: %v"), err)
			}
		}

		return nil
	}

	if f.Config.Period > 0 {
		ticker := time.NewTicker(time.Minute * time.Duration(f.Config.Period))
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				if err := backupFunc(); err != nil {
					return err
				} else {
					log.Println(Colorize(ColorGreen, "Backup completed successfully"))
				}
			case <-ctx.Done():
				log.Println(Colorize(ColorRed, "Backup operation stopped"))
				return ctx.Err()
			}
		}
	} else {
		return backupFunc()
	}
}
