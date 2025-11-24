package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func SelectOne(label string, options []string) (string, error) {
	fmt.Println(label)
	for i, opt := range options {
		fmt.Printf("  %d) %s\n", i+1, opt)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("input error: %w", err)
	}

	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(options) {
		return "", fmt.Errorf("invalid selection")
	}

	return options[choice-1], nil
}
