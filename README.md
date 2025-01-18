# acp

A simple CLI tool that combines common Git commands (add, commit, push) into a single command.

## Usage

### Basic usage (uses current branch name as commit message):
```bash
acp
```
This will:
1. Stage all changes (`git add .`)
2. Commit with branch name as message (`git commit -m "<branch-name>"`)
3. Push to remote, setting upstream if needed (`git push` or `git push --set-upstream origin HEAD`)

### Add specific files, then commit and push:
```bash
acp add file1.txt file2.txt
```
This will:
1. Stage specified files (`git add file1.txt file2.txt`)
2. Commit with branch name as message
3. Push to remote (with upstream handling)

### Add files without committing (staging only):
```bash
acp add file1.txt file2.txt --append
```
This will only stage the specified files (`git add file1.txt file2.txt`)

### Pass-through git commands:
```bash
acp status              # runs: git status
acp checkout -b fix-bug # runs: git checkout -b fix-bug
acp log                 # runs: git log
# etc...
```
Any command other than 'add' will be passed directly to git with all its arguments.