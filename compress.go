package main

import (
	"archive/zip"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func CompressToZip(src, dest string) error {
	zipFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel(filepath.Dir(src), path)
		if err != nil {
			return err
		}

		if info.IsDir() {
			_, err := zipWriter.Create(relativePath + "/")
			return err
		}

		zipFileWriter, err := zipWriter.Create(relativePath)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(zipFileWriter, file)
		return err
	})

	return err
}
