# Code Formatting Setup

This document describes the code formatting and linting setup for the backend of the Reddit clone project.

## Code Formatting Tools

The project uses the following tools for code formatting and linting:

1. **gofmt** - Go's standard code formatter
2. **golangci-lint** - A linter aggregator for Go
3. **pre-commit** - A framework for managing git pre-commit hooks

## Automatic Formatting with Pre-commit Hooks

### Option 1: Using Git's Native Pre-commit Hook

A pre-commit hook is installed in `.git/hooks/pre-commit`. This hook:
- Checks if Go files are properly formatted using `gofmt`
- Automatically formats any unformatted files and adds them to your commit
- Runs `golangci-lint` to check for code quality issues

The pre-commit hook is automatically executed when you run `git commit`.

### Option 2: Using Pre-commit Framework

For a more robust solution, you can use the pre-commit framework:

1. Install pre-commit:
   ```bash
   pip install pre-commit
   ```

2. Install the git hooks:
   ```bash
   pre-commit install
   ```

3. Run against all files:
   ```bash
   pre-commit run --all-files
   ```

The configuration is stored in `.pre-commit-config.yaml` at the root of the repository.

## GitHub Actions

A GitHub Actions workflow is configured to check code formatting and run linters on every pull request and push to the main and dev branches. The workflow is defined in `.github/workflows/go-format.yml`.

## Manual Code Formatting

You can manually format your code with:

```bash
# Format all Go files in the current directory and subdirectories
gofmt -w .

# Run linters
golangci-lint run ./...
```

## Configuration

- `.golangci.yml` - Configuration for golangci-lint
- `.pre-commit-config.yaml` - Configuration for pre-commit framework
- `.github/workflows/go-format.yml` - GitHub Actions workflow configuration 