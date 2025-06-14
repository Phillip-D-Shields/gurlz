/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"gurlz/internal"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	storageManager *internal.StorageManager

	// Lipgloss styles
	rootTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6B9D")).
			Background(lipgloss.Color("#1A1A2E")).
			Width(55).
			Padding(1, 2).
			Align(lipgloss.Center)

	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FF6B9D")).
			Background(lipgloss.Color("#1A1A2E")).
			Width(55).
			Align(lipgloss.Center)

	subtitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#9CA3AF")).
			Italic(true).
			MarginBottom(1)

	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true)

	exampleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#60A5FA")).
			Italic(true).
			PaddingLeft(2)

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#F59E0B")).
			Bold(true)
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gurlz",
	Short: "golang url tester",
	Long:  rootStyledLongDescription(),
	Run: func(cmd *cobra.Command, args []string) {
		showWelcome()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	var err error
	storageManager, err = internal.NewStorageManager()
	if err != nil {
		fmt.Printf("Error initializing storage: %v\n", err)
		os.Exit(1)
	}
}

func rootStyledLongDescription() string {
	title := titleStyle.Render("🌐 GURLZ - Golang URLs")
	subtitle := subtitleStyle.Render("A fast, lightweight CLI tool for managing and testing HTTP requests")

	description := lipgloss.JoinVertical(lipgloss.Left,
		title,
		subtitle,
		"",
		"Perfect for developers who need to:",
		exampleStyle.Render("• Save and organize HTTP requests"),
		exampleStyle.Render("• Test APIs quickly from the terminal"),
		exampleStyle.Render("• Manage headers and request bodies"),
		exampleStyle.Render("• Execute requests and view responses"),
		"",
		infoStyle.Render("💡 Get started: gurlz add my-api https://api.example.com"),
	)

	return description
}

func showWelcome() {
	welcome := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF6B9D")).
		MarginTop(1).
		MarginBottom(1).
		Width(55)

	content := lipgloss.JoinVertical(lipgloss.Left,
		rootTitleStyle.Render("Welcome to Gurlz! 🚀"),
		"",
		"Quick commands to get you started:",
		"",
		commandStyle.Render("📝 Add a request:"),
		exampleStyle.Render("gurlz add api-test https://httpbin.org/get"),
		"",
		commandStyle.Render("📋 List requests:"),
		exampleStyle.Render("gurlz list"),
		"",
		commandStyle.Render("🚀 Execute request:"),
		exampleStyle.Render("gurlz ping api-test"),
		"",
		commandStyle.Render("❓ Get help:"),
		exampleStyle.Render("gurlz --help"),
		"",
		subtitleStyle.Render("Your requests are stored in: ~/.gurlz/"),
	)

	fmt.Println(welcome.Render(content))
}
