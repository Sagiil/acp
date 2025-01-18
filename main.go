package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// Step 1: git add .
	addCmd := exec.Command("git", "add", ".")
	if err := addCmd.Run(); err != nil {
		fmt.Println("Error staging changes:", err)
		os.Exit(1)
	}

	// Get current branch name
	branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	branchOutput, err := branchCmd.Output()
	if err != nil {
		fmt.Println("Error getting branch name:", err)
		os.Exit(1)
	}
	branchName := strings.TrimSpace(string(branchOutput))

	// Step 2: git commit with branch name
	commitCmd := exec.Command("git", "commit", "-m", branchName)
	if err := commitCmd.Run(); err != nil {
		fmt.Println("Error committing changes:", err)
		os.Exit(1)
	}

	// Check if branch has upstream
	hasUpstreamCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	if err := hasUpstreamCmd.Run(); err != nil {
		// No upstream exists, push with --set-upstream
		pushCmd := exec.Command("git", "push", "--set-upstream", "origin", "HEAD")
		if err := pushCmd.Run(); err != nil {
			fmt.Println("Error pushing changes:", err)
			os.Exit(1)
		}
	} else {
		// Upstream exists, just push
		pushCmd := exec.Command("git", "push")
		if err := pushCmd.Run(); err != nil {
			fmt.Println("Error pushing changes:", err)
			os.Exit(1)
		}
	}

	fmt.Println("Successfully added, committed, and pushed changes!")
}
