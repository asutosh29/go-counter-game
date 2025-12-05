package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				return SubmitRename(&m, msg)
			case "esc":
				return QuitRename(&m, msg)
			}
		}
		var cmds []tea.Cmd
		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
		m.viewPort, cmd = m.viewPort.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return QuitGame(&m, msg)
		case "e":
			return StartRename(&m, msg)
		case "up", "j":
			Increment(&m, msg)
		case "down", "k":
			Decrement(&m, msg)
		case "tab":
			m.currentPlayer = (m.currentPlayer + 1) % (len(m.players))
		}
	}
	var cmds []tea.Cmd
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	m.viewPort, cmd = m.viewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
