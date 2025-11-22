/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/praneeth-ayla/AutoCommenter/internal/ai"
	"github.com/praneeth-ayla/AutoCommenter/internal/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command, used to scan files and generate comments.
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and list files needing comments",
	RunE: func(cmd *cobra.Command, args []string) error {

		fmt.Println("Scanning project files...")

		files, err := scanner.Scan(".")
		if err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}

		fmt.Println("Total files scanned:", len(files))

		filtered := scanner.FilterFilesNeedingComments(files)

		fmt.Println("Files considered for commenting:", len(filtered))

		if len(filtered) == 0 {
			fmt.Println("No files found that need comments")
			return nil
		}

		batches := BatchByLines(filtered, 300)
		fmt.Println("Batches created:", len(batches))

		t := ai.NewProvider("gemini")

		for i, b := range batches {
			fmt.Println()
			fmt.Println("Starting batch", i+1, "with", len(b), "files")

			data := scanner.Load(b)

			fmt.Println("Loaded content for batch", i+1)

			comments, err := t.GenerateContextBatch(data)
			if err != nil {
				fmt.Println("Error generating comments for batch", i+1, err)
				continue
			}

			fmt.Println("Completed batch", i+1)
			fmt.Println("------------------------------------------------")

			for _, c := range comments {
				fmt.Println("File:", c.Path)
				fmt.Println("Exports:", c.Exports)
				fmt.Println("Imp Logic:", c.ImpLogic)
				fmt.Println("Name:", c.Name)
				fmt.Println("Summary:", c.Summary)
				fmt.Println("------------------------------------------------")
			}
		}

		fmt.Println()
		fmt.Println("Scan completed successfully")

		return nil
	},
}

func init() {
	// Add the scan command as a subcommand to the root command.
	rootCmd.AddCommand(scanCmd)
}

func BatchByLines(files []scanner.Info, maxLines int) [][]scanner.Info {
	var result [][]scanner.Info
	var group []scanner.Info
	used := 0

	for _, f := range files {
		if used+f.Lines > maxLines {
			result = append(result, group)
			group = []scanner.Info{}
			used = 0
		}
		group = append(group, f)
		used += f.Lines
	}

	if len(group) > 0 {
		result = append(result, group)
	}

	return result
}
