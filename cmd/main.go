package main

import (
	"flag"
	"fmt"

	"github.com/cjnghn/gfd/downloader"
	"github.com/cjnghn/gfd/parser"
)

func main() {
	urlFlag := flag.String("url", "", "GitHub URL of the folder to download")
	outputDir := flag.String("output", "", "Output directory for the downloaded contents")

	flag.Parse()

	if *urlFlag == "" || *outputDir == "" {
		flag.Usage()
		return
	}

	owner, repo, path, err := parser.ParseGitHubURL(*urlFlag)
	if err != nil {
		fmt.Printf("Error parsing URL: %v\n", err)
		return
	}

	if err := downloader.DownloadFolder(owner, repo, path, *outputDir); err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		return
	}

	fmt.Println("Download completed successfully.")
}
