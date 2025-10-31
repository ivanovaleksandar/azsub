package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// isFzfAvailable checks if fzf binary exists in PATH
func isFzfAvailable() bool {
	// Allow users to disable fzf with environment variable
	if os.Getenv("AZSUB_IGNORE_FZF") == "1" {
		return false
	}

	_, err := exec.LookPath("fzf")
	return err == nil
}

// selectWithFzf pipes subscriptions to fzf and returns selected subscription
func selectWithFzf(subscriptions []Subscription) (*Subscription, error) {
	// Prepare input data (tab-separated: Name<TAB>ID)
	var input strings.Builder
	for _, sub := range subscriptions {
		fmt.Fprintf(&input, "%s\t%s\n", sub.Name, sub.ID)
	}

	// Create fzf command with options
	cmd := exec.Command("fzf",
		"--height=50%",
		"--reverse",
		"--header=Select Azure Subscription",
		"--delimiter=\t",
		"--with-nth=1", // Only search on name column
	)

	// Pipe subscription data to fzf
	cmd.Stdin = strings.NewReader(input.String())
	cmd.Stderr = os.Stderr

	// Capture selected line
	output, err := cmd.Output()
	if err != nil {
		return nil, err // User cancelled or fzf error
	}

	// Parse selection (extract ID from tab-separated line)
	parts := strings.Split(strings.TrimSpace(string(output)), "\t")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid selection format")
	}
	selectedID := parts[1]

	// Find and return the full subscription object
	for i := range subscriptions {
		if subscriptions[i].ID == selectedID {
			return &subscriptions[i], nil
		}
	}

	return nil, fmt.Errorf("selected subscription not found")
}
