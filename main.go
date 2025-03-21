package main

import (
	"fmt"
	"github-user-activity/handlers"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}

	username := os.Args[1]
	guf, err := handlers.NewGithubUserFetcher(username)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	guf.DisplayEvents()

}

func printUsage() {
	fmt.Println("Usage: github-activity <username>")
	os.Exit(1)
}
