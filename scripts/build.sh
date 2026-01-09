#!/bin/bash
set -e

echo "🚀 Building Anton TUI..."

mkdir -p bin

# On demande à Go de compiler ce qui se trouve dans le dossier ./cmd/
go build -o bin/anton-tui ./cmd/main.go

echo "✅ Build terminé !"