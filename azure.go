package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// Subscription represents an Azure subscription
type Subscription struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	State     string `json:"state"`
	IsDefault bool   `json:"isDefault"`
	TenantID  string `json:"tenantId"`
}

// listSubscriptions retrieves all Azure subscriptions using az CLI
func listSubscriptions() ([]Subscription, error) {
	cmd := exec.Command("az", "account", "list", "--output", "json")
	output, err := cmd.Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return nil, fmt.Errorf("az CLI error: %s", string(exitErr.Stderr))
		}
		return nil, fmt.Errorf("failed to execute az CLI: %w", err)
	}

	var subscriptions []Subscription
	if err := json.Unmarshal(output, &subscriptions); err != nil {
		return nil, fmt.Errorf("failed to parse subscription list: %w", err)
	}

	return subscriptions, nil
}

// setSubscription sets the active Azure subscription
func setSubscription(subscriptionID string) error {
	cmd := exec.Command("az", "account", "set", "--subscription", subscriptionID)
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("az CLI error: %s", string(exitErr.Stderr))
		}
		return fmt.Errorf("failed to set subscription: %w", err)
	}

	return nil
}
