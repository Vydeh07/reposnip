package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	token := flag.String("token", "", "GitHub token for private repos")
	output := flag.String("output", ".", "Folder to save files into")

	flag.Parse()

	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: reposnip <github-url> [--token your_token]")
		os.Exit(1)
	}

	url := args[0]

	files, err := FetchFiles(url, *token)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\nFound %d files:\n", len(files))
	for _, f := range files {
		fmt.Println(" ", f.Path)
	}

	fmt.Println()
	err = DownloadFiles(files, *token, *output)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("\n Done! %d file(s) downloaded.\n", len(files))
}
