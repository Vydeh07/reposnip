package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func DownloadFiles(files []GithubFile, token string, output string) error {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var firstErr error

	fmt.Printf("Downloading %d files concurrently...\n", len(files))

	for _, file := range files {
		wg.Add(1)
		go func(f GithubFile) {
			defer wg.Done()
			err := downloadOne(f, token, output)
			if err != nil {
				mu.Lock()
				if firstErr == nil {
					firstErr = fmt.Errorf("failed to download %s: %w", f.Path, err)
				}
				mu.Unlock()
				return
			}
			mu.Lock()
			fmt.Println("  ↓", f.Path)
			mu.Unlock()

		}(file)
	}
	wg.Wait()
	return firstErr
}

func downloadOne(file GithubFile, token string, output string) error {
	body, err := makeRequest(file.DownloadURL, token)
	if err != nil {
		return err
	}
	fullPath := filepath.Join(output, file.Path)
	dir := filepath.Dir(fullPath)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}
	return os.WriteFile(fullPath, body, 0644)
}
