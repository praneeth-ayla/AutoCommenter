/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/praneeth-ayla/AutoCommenter/internal/ai"
	"github.com/praneeth-ayla/AutoCommenter/internal/contextstore"
	"github.com/praneeth-ayla/AutoCommenter/internal/scanner"
	"github.com/spf13/cobra"
)

// readmeCmd represents the readme command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("readme called")
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)
	readmeCmd.AddCommand(genReadmeCmd)
}

var genReadmeCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate README.md for the project",
	RunE: func(cmd *cobra.Command, args []string) error {

		rootPath := scanner.GetProjectRoot()
		provider := ai.NewProvider("gemini")

		fmt.Println("Loading project context...")
		contextData, err := contextstore.Load()
		if err != nil {
			return fmt.Errorf("failed to load context: %w", err)
		}
		allCtxSlice := contextstore.MapToSlice(contextData)

		// check existing README file
		var existingReadme string
		readmePaths := []string{
			filepath.Join(rootPath, "README.md"),
			filepath.Join(rootPath, "readme.md"),
		}

		for _, path := range readmePaths {
			if data, err := os.ReadFile(path); err == nil {
				existingReadme = string(data)
				fmt.Println("Existing README found:", path)
				break
			}
		}

		fmt.Println("️Generating README...")
		newReadme, err := provider.GenerateReadme(allCtxSlice, existingReadme)
		if err != nil {
			return fmt.Errorf("README generation failed: %w", err)
		}

		outputPath := filepath.Join(rootPath, "README.md")

		if err := os.WriteFile(outputPath, []byte(newReadme), 0644); err != nil {
			return fmt.Errorf("failed to write README.md: %w", err)
		}

		fmt.Println(";) README.md updated:", outputPath)
		return nil
	},
}
