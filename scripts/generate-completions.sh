#!/bin/bash
set -e

# Create completions directory
mkdir -p completions

# Build the binary temporarily to generate completions
go build -o bm ./main.go

# Generate shell completions
./bm completion bash > completions/bm.bash
./bm completion zsh > completions/bm.zsh
./bm completion fish > completions/bm.fish

# Clean up temporary binary
rm bm