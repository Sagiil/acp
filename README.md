# acp

A simple CLI tool that combines common Git commands (add, commit, push) into a single command.

## Usage

Basic usage (uses current branch name as commit message):

```bash
acp
```

## What it does

1. Stages all changes (`git add .`)
2. Commits with branch name as message (`git commit -m "<branch-name>"`)
3. Pushes to remote, setting upstream if needed (`git push` or `git push --set-upstream origin HEAD`)