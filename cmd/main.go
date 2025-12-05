package main

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Styles
const (
	primaryColor   = "#B8BB26"   // A nice lime green
	secondaryColor = "#D3869B"   // A soft purple
	subtleColor    = "#504945"   // Dark grey
	dangerColor    = "#cc4949ff" //Danger color
)

var (
	// Style for the "Current Count" label
	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(secondaryColor)).
			Bold(true).
			MarginBottom(1) // Add space below the label

	// Style for the actual number
	counterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(primaryColor)).
			Background(lipgloss.Color(subtleColor)).
			Padding(1, 3). // Top/Bottom: 1, Left/Right: 3
			Align(lipgloss.Center).
			Bold(true)

	// Style for the help text at the bottom
	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")). // ANSI 256 color code for grey
			Italic(true).
			MarginTop(2) // Add space above the help text

	// Style for the help text at the bottom
	messageStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(dangerColor)). // ANSI 256 color code for grey
			Italic(true).
			MarginTop(2) // Add space above the help text

	selectedContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(primaryColor))
	containerStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))
)

type player struct {
	counter int
	name    string
}

type model struct {
	players       []player
	message       string
	currentPlayer int // will 0 or 1

	textInput textinput.Model
	isEditing bool
}

func (m model) Init() tea.Cmd { return nil }

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Enter new name"
	ti.CharLimit = 20
	ti.Width = 30

	return model{
		players: []player{
			{
				name: "Player 1",
			},
			{
				name: "Player 2",
			},
		},
		textInput: ti,
		isEditing: false,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if m.isEditing {
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "enter":
				if m.textInput.Value() != "" {
					m.players[m.currentPlayer].name = m.textInput.Value()
				}

				m.isEditing = false
				m.textInput.Blur()
				return m, nil
			case "esc":
				m.isEditing = false
				m.textInput.Blur()
				return m, nil
			}
		}

		var cmd tea.Cmd
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "e":
			m.isEditing = true
			m.textInput.Focus()
			m.textInput.SetValue(m.players[m.currentPlayer].name)
			return m, textinput.Blink

		case "up", "j":
			if m.players[m.currentPlayer].counter >= 10 {
				m.message = "Can't go above 10..."
				return m, nil
			}
			m.message = ""
			m.players[m.currentPlayer].counter++
		case "down", "k":
			if m.players[m.currentPlayer].counter <= 0 {
				m.message = "Can't go below 0..."
				return m, nil
			}
			m.message = ""
			m.players[m.currentPlayer].counter--
		case "tab":
			m.currentPlayer = (m.currentPlayer + 1) % len(m.players)
		}
	}
	return m, nil
}
func (m model) View() string {

	label := labelStyle.Render("=== Counter Game ===")

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
		// Render the bubble component
		footer = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(secondaryColor)).
			Padding(0, 1).
			Render(m.textInput.View())
	} else {
		// Render your old message/help text
		footer = messageStyle.Render(m.message) + "\n" + helpStyle.Render("• Press up/down or j/k to increase/decrease the counter \n• Press 'e' to rename \n• Tab to switch")
	}

	ui := lipgloss.JoinVertical(lipgloss.Center, label, playersRowString, footer)

	container := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()). // Nice rounded corners
		BorderForeground(lipgloss.Color(primaryColor)).
		Padding(1, 2). // Add breathing room inside the border
		Render(ui)
	return container
}

func main() {
	program := tea.NewProgram(initialModel())
	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}
