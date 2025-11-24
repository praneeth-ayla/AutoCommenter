/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/praneeth-ayla/AutoCommenter/internal/config"
	"github.com/spf13/cobra"
)

var providerCmd = &cobra.Command{
	Use:   "provider",
	Short: "Manage AI provider setting",
	Long:  `Set or view which AI provider is used for context generation.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'AutoCommenter provider set' to interactively set provider")
	},
}

// List of supported providers
var supportedProviders = []string{
	"gemini",
}

var providerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Interactively select an AI provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Select AI Provider:")
		for i, p := range supportedProviders {
			fmt.Printf("  %d) %s\n", i+1, p)
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter choice number: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("input error: %w", err)
		}

		input = strings.TrimSpace(input)
		choice, err := strconv.Atoi(input)
		if err != nil || choice < 1 || choice > len(supportedProviders) {
			return fmt.Errorf("invalid selection")
		}

		selected := supportedProviders[choice-1]
		if err := config.SetProvider(selected); err != nil {
			return fmt.Errorf("could not save provider: %w", err)
		}

		fmt.Println("Provider updated to", selected)
		return nil
	},
}

var providerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Show current AI provider",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := config.GetProvider()
		fmt.Println("Current provider:", name)
	},
}

func init() {
	rootCmd.AddCommand(providerCmd)
	providerCmd.AddCommand(providerSetCmd)
	providerCmd.AddCommand(providerGetCmd)
}
