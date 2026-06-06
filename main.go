package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("PR review bot started")
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
	fmt.Printf("Reviewing pull request #%d\n", event.Number)
}
