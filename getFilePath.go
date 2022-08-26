package main

import (
	"log"
	"os"
	"path/filepath"
)

func getFilePath(dirpath string) []string {
	var paths []string

	// func ReadDir(name string) ([]DirEntry, error)
	files, err := os.ReadDir(dirpath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		// ディレクトリの場合
		if file.IsDir() {
			paths = append(paths, getFilePath(filepath.Join(dirpath, file.Name()))...)
		}
		if !file.IsDir() {
			paths = append(paths, filepath.Join(dirpath, file.Name()))
		}
	}

	return paths
}
