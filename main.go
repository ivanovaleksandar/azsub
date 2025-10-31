package main

import (
	"fmt"
	"os"
)

func main() {
	// Get all Azure subscriptions
	subscriptions, err := listSubscriptions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to list subscriptions: %v\n", err)
		os.Exit(1)
	}

	if len(subscriptions) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no subscriptions found\n")
		os.Exit(1)
	}

	// Show interactive fuzzy finder (or list if fzf not available)
	selected, err := selectSubscription(subscriptions)
	if err != nil {
		// If fzf is not available, list was already printed
		// Exit cleanly without error code
		os.Exit(0)
	}

	// Set the Azure account
	if err := setSubscription(selected.ID); err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to set subscription: %v\n", err)
		os.Exit(1)
	}

	// Output shell commands to set environment variables
	fmt.Printf("export ARM_SUBSCRIPTION_ID=\"%s\"\n", selected.ID)
	fmt.Printf("export ARM_SUBSCRIPTION_NAME=\"%s\"\n", selected.Name)
	fmt.Fprintf(os.Stderr, "✓ Switched to subscription: %s\n", selected.Name)
}
