#!/bin/bash

# Couleurs pour les messages
RED='\033[0;31m'
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m'

print_banner() {
    echo -e "${RED}"
    cat << "EOF"
╔═══════════════════════════════════════════╗
║             ANTON - Terminal UI           ║
║     "Efficiency through Simplicity"       ║
╚═══════════════════════════════════════════╝
EOF
    echo -e "${NC}"
}

print_message() {
    echo -e "${BLUE}[ANTON]${NC} $1"
}

print_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

print_error() {
    echo -e "${RED}[ERROR]${NC} $1"
    exit 1
}

# Afficher la bannière
print_banner

# Vérifier Go
if ! command -v go &> /dev/null; then
    print_error "Go n'est pas installé. Installation requise."
    exit 1
fi

# Nettoyer
print_message "Nettoyage de l'environnement..."
rm -rf bin/
mkdir -p bin/

# Dépendances
print_message "Vérification des dépendances..."
go mod tidy || print_error "Erreur lors de la vérification des dépendances"

# Installation des dépendances nécessaires
print_message "Installation des dépendances..."
go get github.com/charmbracelet/bubbletea
go get github.com/charmbracelet/bubbles
go get github.com/charmbracelet/lipgloss

# Build
print_message "Construction du projet..."
env CGO_ENABLED=0 go build -o bin/anton-tui cmd/main.go || print_error "Erreur lors de la compilation"

chmod +x bin/anton-tui
print_success "Build terminé avec succès."

# Proposition de lancement
read -p "Voulez-vous lancer anton-tui maintenant ? (o/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Oo]$ ]]; then
    print_message "Lancement d'Anton..."
    ./bin/anton-tui
fi