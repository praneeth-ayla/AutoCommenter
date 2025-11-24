/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/praneeth-ayla/AutoCommenter/internal/ai"
	"github.com/praneeth-ayla/AutoCommenter/internal/config"
	"github.com/praneeth-ayla/AutoCommenter/internal/ui"
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

var providerSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Interactively select an AI provider",
	RunE: func(cmd *cobra.Command, args []string) error {
		selectedProvider, err := ui.SelectOne("Select AI Provider:", ai.SupportedProviders)
		if err != nil {
			return err
		}

		err = config.SetProvider(selectedProvider)
		if err != nil {
			return fmt.Errorf("could not save provider: %w", err)
		}

		fmt.Println("Provider updated to", selectedProvider)
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
