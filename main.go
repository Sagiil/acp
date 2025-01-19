package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func gitAdd(files []string) error {
	args := append([]string{"add"}, files...)
	addCmd := exec.Command("git", args...)
	return addCmd.Run()
}


func gitCommitAndPush(files []string, commitMessage string) error {
	// If no commit message provided, use branch name
	if commitMessage == "" {
		// Get current branch name
		branchCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
		branchOutput, err := branchCmd.Output()
		if err != nil {
			return fmt.Errorf("error getting branch name: %w", err)
		}
		commitMessage = strings.TrimSpace(string(branchOutput))
	}

	// Commit
	commitCmd := exec.Command("git", "commit", "-m", commitMessage)
	if err := commitCmd.Run(); err != nil {
		return fmt.Errorf("error committing changes: %w", err)
	}

	// Check if branch has upstream
	hasUpstreamCmd := exec.Command("git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}")
	if err := hasUpstreamCmd.Run(); err != nil {
		// No upstream exists, push with --set-upstream
		pushCmd := exec.Command("git", "push", "--set-upstream", "origin", "HEAD")
		if err := pushCmd.Run(); err != nil {
			return fmt.Errorf("error pushing changes: %w", err)
		}
	} else {
		// Upstream exists, just push
		pushCmd := exec.Command("git", "push")
		if err := pushCmd.Run(); err != nil {
			return fmt.Errorf("error pushing changes: %w", err)
		}
	}

	return nil
}

func executeGitCommand(args []string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	args := os.Args[1:]

	// Parse for commit message flag
	var commitMessage string
	var newArgs []string
	for i := 0; i < len(args); i++ {
		if args[i] == "--m" {
			if i+1 < len(args) {
				commitMessage = args[i+1]
				i++ // skip next arg since it's the message
				continue
			}
		}
		newArgs = append(newArgs, args[i])
	}
	args = newArgs

	// If no arguments, perform default behavior (add all, commit, push)
	if len(args) == 0 {
		if err := gitAdd([]string{"."}); err != nil {
			fmt.Println("Error staging changes:", err)
			os.Exit(1)
		}
		if err := gitCommitAndPush([]string{"."}, commitMessage); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("Successfully added, committed, and pushed all changes!")
		return
	}

	// Handle "add" command specially
	if args[0] == "add" {
		if len(args) < 2 {
			fmt.Println("Error: 'add' command requires at least one file")
			os.Exit(1)
		}

		// Check if --append flag is present
		isAppend := false
		files := []string{}
		for _, arg := range args[1:] {
			if arg == "--append" {
				isAppend = true
			} else {
				files = append(files, arg)
			}
		}

		// Add the specified files
		if err := gitAdd(files); err != nil {
			fmt.Println("Error staging files:", err)
			os.Exit(1)
		}

		if !isAppend {
			// If not append, also commit and push
			if err := gitCommitAndPush(files, commitMessage); err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("Successfully added, committed, and pushed files: %v\n", files)
		} else {
			fmt.Printf("Successfully added files: %v\n", files)
		}
		return
	}

	// For all other commands, pass directly to git
	if err := executeGitCommand(args); err != nil {
		fmt.Println("Error executing git command:", err)
		os.Exit(1)
	}
}