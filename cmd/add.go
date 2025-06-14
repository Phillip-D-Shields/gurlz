/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"gurlz/internal"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var (
	method  string
	headers []string
	body    string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new request",
	Long: `Add a new HTTP request to your collection.
		
Examples:
  gurlz add api-health https://api.example.com/health
  gurlz add api-users https://api.example.com/users -X POST -H "Content-Type: application/json" -d '{"name":"test"}'
  gurlz add api-auth https://api.example.com/login -X POST -H "Authorization: Bearer token123"`,
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		url := args[1]

		// check name
		if strings.Contains(name, " ") {
			return fmt.Errorf("request name cannot contain spaces")
		}
		if len(name) > 50 {
			return fmt.Errorf("request name must be 50 characters or less")
		}

		// parse headers
		headerMap := make(map[string]string)
		for _, header := range headers {
			parts := strings.SplitN(header, ":", 2)
			if len(parts) != 2 {
				return fmt.Errorf("invalid header format: %s (expected 'Key: Value')", header)
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headerMap[key] = value
		}

		// new request
		now := time.Now()
		req := internal.Request{
			ID:        uuid.New().String(),
			Name:      name,
			URL:       url,
			Method:    strings.ToUpper(method),
			Headers:   headerMap,
			Body:      body,
			CreatedAt: now,
			UpdatedAt: now,
		}
		store, err := storageManager.LoadRequests()
		if err != nil {
			return fmt.Errorf("failed to load requests: %w", err)
		}

		// Add request
		if err := store.AddRequest(req); err != nil {
			return err
		}

		// Save to file
		if err := storageManager.SaveRequests(store); err != nil {
			return fmt.Errorf("failed to save request: %w", err)
		}

		fmt.Printf("✅ Added request '%s' -> %s %s\n", name, method, url)
		if len(headerMap) > 0 {
			fmt.Printf("   Headers: %d\n", len(headerMap))
		}
		if body != "" {
			fmt.Printf("   Body: %d characters\n", len(body))
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Flags
	addCmd.Flags().StringVarP(&method, "method", "X", "GET", "HTTP method")
	addCmd.Flags().StringArrayVarP(&headers, "header", "H", []string{}, "HTTP headers (can be used multiple times)")
	addCmd.Flags().StringVarP(&body, "data", "d", "", "Request body data")
}
