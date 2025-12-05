package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {

	label := labelStyle.Render("=== Counter Game ===")

	timerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("205")) // Pink
	if m.timeLeft < 5 {
		timerStyle = timerStyle.Foreground(lipgloss.Color("#FF0000")).Bold(true)
	}
	timerView := timerStyle.Render(fmt.Sprintf("Time Left: %ds", m.timeLeft))

	playersRow := []string{}
	for i, player := range m.players {
		name := labelStyle.Render(player.name)
		counter := counterStyle.Render(fmt.Sprintf("Score: %d", player.counter))

		container := lipgloss.JoinVertical(lipgloss.Center, name, counter)
		if m.currentPlayer == i {
			container = selectedContainerStyle.Render(container)
		} else {
			container = containerStyle.Render(container)
		}
		playersRow = append(playersRow, container)
	}
	playersRowString := lipgloss.JoinHorizontal(lipgloss.Center, playersRow...)

	var footer string
	if m.isEditing {
		footer = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(secondaryColor)).
			Padding(0, 1).
			Render(m.textInput.View())
	} else {
		footer = helpStyle.Render("• Press up/down or j/k to increase/decrease the counter \n• Press 'e' to rename \n• Tab to switch")
	}

	logString := ""
	for _, log := range m.logMessages {
		logString += log + "\n"
	}
	m.viewPort.SetContent(messageStyle.Render(logString + "\n" + m.message))
	logView := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(subtleColor)).
		Padding(0, 1).
		Render(m.viewPort.View())

	ui := lipgloss.JoinVertical(lipgloss.Center, label, timerView, playersRowString, footer, logView)

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()). // Nice rounded corners
		BorderForeground(lipgloss.Color(primaryColor)).
		Padding(1, 2). // Add breathing room inside the border
		Render(ui)
	return container
}
