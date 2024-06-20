package downloader

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const GITHUB_API_URL = "https://api.github.com/repos"

type GitHubContent struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Type        string `json:"type"`
	DownloadURL string `json:"download_url"`
}

func downloadFile(url string, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadFolder(owner, repo, path, outputDir string) error {
	apiURL := fmt.Sprintf("%s/%s/%s/contents/%s", GITHUB_API_URL, owner, repo, path)
	resp, err := http.Get(apiURL)
	if err != nil {
		return fmt.Errorf("failed to fetch repository contents: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("failed to fetch repository contents: %s", resp.Status)
	}

	var contents []GitHubContent
	if err := json.NewDecoder(resp.Body).Decode(&contents); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for _, item := range contents {
		itemPath := filepath.Join(outputDir, item.Name)
		if item.Type == "file" {
			fmt.Printf("Downloading file: %s\n", item.DownloadURL)
			if err := downloadFile(item.DownloadURL, itemPath); err != nil {
				return err
			}
		} else if item.Type == "dir" {
			fmt.Printf("Entering directory: %s\n", item.Path)
			if err := DownloadFolder(owner, repo, item.Path, itemPath); err != nil {
				return err
			}
		}
	}

	return nil
}
