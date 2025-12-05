package main

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Rename Actions
func StartRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.isEditing = true
	m.textInput.Focus()
	m.textInput.SetValue(m.players[m.currentPlayer].name)

	m.message = "Editting name... press esc to exit!"
	m.logMessages = append(m.logMessages, "Editting name")
	return *m, textinput.Blink
}
func SubmitRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.textInput.Value() != "" {
		m.players[m.currentPlayer].name = m.textInput.Value()
	}

	m.isEditing = false
	m.textInput.Blur()
	return *m, nil
}
func QuitRename(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.isEditing = false
	m.textInput.Blur()
	return *m, nil
}

// Game actions
func Increment(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.players[m.currentPlayer].counter < 10 {
		m.players[m.currentPlayer].counter++
		m.message = ""

		logEntry := fmt.Sprintf("%s scored! Total: %d", m.players[m.currentPlayer].name, m.players[m.currentPlayer].counter)
		m.logMessages = append(m.logMessages, logEntry)

		m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
		m.viewPort.GotoBottom()
	} else {
		m.message = "Can't go above 10..."
	}
	return *m, nil
}
func Decrement(m *model, msg tea.Msg) (model, tea.Cmd) {
	if m.players[m.currentPlayer].counter > 0 {
		m.players[m.currentPlayer].counter--
		m.message = ""

		logEntry := fmt.Sprintf("%s lost a point. Total: %d", m.players[m.currentPlayer].name, m.players[m.currentPlayer].counter)
		m.logMessages = append(m.logMessages, logEntry)

		m.viewPort.SetContent(strings.Join(m.logMessages, "\n"))
		m.viewPort.GotoBottom()
	} else {
		m.message = "Can't go below 0..."
	}
	return *m, nil
}

func QuitGame(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.message = "Quitting thanks for playing!"
	return *m, tea.Quit
}

// Timer Actions
func UpdateTimer(m *model, msg tea.Msg) (model, tea.Cmd) {
	m.timeLeft--
	return *m, tick()
}
