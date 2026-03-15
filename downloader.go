package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func DownloadFiles(files []GithubFile, token string, output string) error {
	for _, file := range files {
		err := downloadOne(file, token, output)
		if err != nil {
			return fmt.Errorf("failed to download %s: %w", file.Path, err)
		}
		fmt.Println("downloaded:", file.Path)
	}
	return nil
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
