package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type GithubFile struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

func FetchFiles(rawURL string, token string) ([]GithubFile, error) {

	owner, repo, branch, path, err := parseGitHubURL(rawURL)
	if err != nil {
		return nil, err
	}
	return fetchDir(owner, repo, branch, path, token)
}
func fetchDir(owner, repo, branch, path, token string) ([]GithubFile, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/%s/contents/%s?ref=%s", owner, repo, path, branch)
	body, err := makeRequest(apiURL, token)
	if err != nil {
		return nil, err
	}
	var items []GithubFile
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	var allFiles []GithubFile
	for _, item := range items {
		if item.Type == "file" {
			allFiles = append(allFiles, item)
		} else if item.Type == "dir" {
			subFiles, err := fetchDir(owner, repo, branch, item.Path, token)
			if err != nil {
				return nil, err
			}
			allFiles = append(allFiles, subFiles...)
		}
	}

	return allFiles, nil
}
func makeRequest(url, token string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "graft-cli")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("GitHub API error: status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}

func parseGitHubURL(rawURL string) (string, string, string, string, error) {
	rawURL = strings.TrimRight(rawURL, "/")
	rawURL = strings.TrimPrefix(rawURL, "https://")
	rawURL = strings.TrimPrefix(rawURL, "http://")

	if !strings.HasPrefix(rawURL, "github.com/") {
		return "", "", "", "", fmt.Errorf("not a GitHub URL: %s", rawURL)
	}
	rawURL = strings.TrimPrefix(rawURL, "github.com/")
	parts := strings.SplitN(rawURL, "/", 3)
	if len(parts) < 2 {
		return "", "", "", "", fmt.Errorf("URL must have at least owner/repo")
	}

	owner := parts[0]
	repo := parts[1]
	if len(parts) < 3 {
		return owner, repo, "main", "", nil
	}
	remainder := parts[2]
	remainder = strings.TrimPrefix(remainder, "tree/")
	remainder = strings.TrimPrefix(remainder, "blob/")
	branchAndPath := strings.SplitN(remainder, "/", 2)
	branch := branchAndPath[0]

	path := ""
	if len(branchAndPath) > 1 {
		path = branchAndPath[1]
	}

	return owner, repo, branch, path, nil
}
