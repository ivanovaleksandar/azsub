package main

import (
	"fmt"
)

// selectSubscription orchestrates selection logic based on fzf availability
func selectSubscription(subscriptions []Subscription) (*Subscription, error) {
	if isFzfAvailable() {
		return selectWithFzf(subscriptions)
	}

	// Fallback: print list and return error indicating fzf is not available
	printSubscriptionList(subscriptions)
	return nil, fmt.Errorf("fzf not found - install fzf for interactive selection")
}

// printSubscriptionList outputs subscriptions in readable format
func printSubscriptionList(subscriptions []Subscription) {
	for _, sub := range subscriptions {
		fmt.Printf("%-50s %s\n", sub.Name, sub.ID)
	}
}
