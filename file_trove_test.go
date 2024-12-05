package main

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func createTestFile(t *testing.T, dir, fileName, content string) string {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create directory '%s': %v", dir, err)
	}

	filePath := filepath.Join(dir, fileName)
	err = os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create file '%s': %v", filePath, err)
	}

	info, err := os.Stat(dir)
	if err != nil {
		log.Fatalf("Failed to stat directory: %v", err)
	}

	if info.Mode().Perm() != 0755 {
		os.Chmod(dir, 0755)
	}

	return filePath
}

func TestBackup(t *testing.T) {
	sourceDir := t.TempDir()
	destDir := t.TempDir()

	config := Config{
		Source:      sourceDir,
		Destination: destDir,
		Compress:    false,
		Period:      0,
	}
	worker := NewFileWorker(config)

	ctx := context.Background()
	err := worker.Backup(ctx)
	if err != nil {
		t.Fatalf("Backup failed: %v", err)
	}

	copiedFile := filepath.Join(destDir, filepath.Base(sourceDir))

	if _, err := os.Stat(copiedFile); err != nil {
		t.Fatalf("Copied file '%s' does not exist or is inaccessible: %v", copiedFile, err)
	}
}

func TestCompressToZip(t *testing.T) {
	sourceDir := t.TempDir()
	createTestFile(t, sourceDir, "file1.txt", "Content 1")
	createTestFile(t, sourceDir, "file2.txt", "Content 2")

	zipPath := filepath.Join(t.TempDir(), "test.zip")
	err := CompressToZip(sourceDir, zipPath)
	if err != nil {
		t.Fatalf("Compression failed: %v", err)
	}

	if _, err := os.Stat(zipPath); os.IsNotExist(err) {
		t.Errorf("ZIP file was not created")
	}
}

func TestLogging(t *testing.T) {
    var buf bytes.Buffer
    log.SetOutput(&buf)

    log.Println("Test log message")

    if !strings.Contains(buf.String(), "Test log message") {
        t.Errorf("Log output does not contain expected message")
    }
}