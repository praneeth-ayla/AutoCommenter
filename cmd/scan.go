/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/praneeth-ayla/AutoCommenter/internal/scanner"
	"github.com/spf13/cobra"
)

// scanCmd represents the scan command
var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "A brief description of your command",
	Long:  `It scans all the files and directory inside`,
	Run: func(cmd *cobra.Command, args []string) {
		files, err := scanner.Scanner(".")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(files)
	},
}

func init() {
	rootCmd.AddCommand(scanCmd)
}
