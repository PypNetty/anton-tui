package main

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	currentTab  int
	tabs        []string
	content     string
	cpuUsage    float64
	memoryUsage float64
	diskUsage   float64
}

func initialModel() model {
	return model{
		tabs:        []string{"System", "Processes", "Network", "Logs"},
		cpuUsage:    0.0,
		memoryUsage: 0.0,
		diskUsage:   0.0,
	}
}

func (m model) Init() tea.Cmd {
	// Commande initiale
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "tab", "right":
			m.currentTab = (m.currentTab + 1) % len(m.tabs)
		case "shift+tab", "left":
			m.currentTab--
			if m.currentTab < 0 {
				m.currentTab = len(m.tabs) - 1
			}
		}
	}
	return m, nil
}

func (m model) View() string {
	// Styles
	tabStyle := lipgloss.NewStyle().
		Padding(0, 2).
		Bold(true)
	activeTabStyle := tabStyle.Copy().
		Foreground(lipgloss.Color("#FFA500")).
		Background(lipgloss.Color("#1A1A1A"))
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#666666")).
		Padding(1).
		Width(50)

	// Rendu des onglets
	var tabs []string
	for i, tab := range m.tabs {
		if i == m.currentTab {
			tabs = append(tabs, activeTabStyle.Render(tab))
		} else {
			tabs = append(tabs, tabStyle.Render(tab))
		}
	}

	// Contenu selon l'onglet actif
	var content string
	switch m.currentTab {
	case 0: // System
		content = boxStyle.Render(fmt.Sprintf(
			"CPU Usage: %.1f%%\nMemory Usage: %.1f%%\nDisk Usage: %.1f%%",
			m.cpuUsage,
			m.memoryUsage,
			m.diskUsage,
		))
	case 1: // Processes
		content = boxStyle.Render("Process list coming soon...")
	case 2: // Network
		content = boxStyle.Render("Network stats coming soon...")
	case 3: // Logs
		content = boxStyle.Render("System logs coming soon...")
	}

	return lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.JoinHorizontal(lipgloss.Left, tabs...),
		"",
		content,
		"",
		"Press q to quit • tab/←→ to switch tabs",
	)
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
