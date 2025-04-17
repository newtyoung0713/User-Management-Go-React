#!/bin/bash

# Install goimports for formatting
go install golang.org/x/tools/cmd/goimports@latest

# Install golangci-lint for linting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install gopls for language server
go install golang.org/x/tools/gopls@latest

echo "Go development tools installed successfully!"
