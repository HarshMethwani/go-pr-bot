package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/google/go-github/v69/github"
)

func main() {
	fmt.Println("PR review bot started")
	prNumber := readPrNumber()
	owner, repo := readRepo()
	files := fetchFilesChanged(owner, repo, prNumber)
	for _, f := range files {
		fmt.Printf("%s ( +%d / -%d) \n", f.GetFilename(), f.GetAdditions(), f.GetDeletions())
	}
}

func readRepo() (string, string) {
	full := os.Getenv("GITHUB_REPOSITORY")
	parts := strings.SplitN(full, "/", 2)
	return parts[0], parts[1]
}
func fetchFilesChanged(owner, repo string, prNumber int) []*github.CommitFile {
	authToken := os.Getenv("GITHUB_TOKEN")
	client := github.NewClient(nil).WithAuthToken(authToken)
	ctx := context.Background()
	files, _, err := client.PullRequests.ListFiles(ctx, owner, repo, prNumber, nil)
	if err != nil {
		fmt.Println("could not fetch files", err)
		os.Exit(1)
	}
	return files
}
func readPrNumber() int {
	eventPath := os.Getenv("GITHUB_EVENT_PATH")
	data, err := os.ReadFile(eventPath)
	if err != nil {
		fmt.Println("Could not read the file")
		os.Exit(1)
	}
	var event struct {
		Number int `json:"number"`
	}
	if err := json.Unmarshal(data, &event); err != nil {
		fmt.Println("could not parse event")
		os.Exit(1)
	}
	return event.Number
}
