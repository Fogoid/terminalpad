package models

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

const (
	username = 0
	password = 1
	header   = " _____                   _             _ ____           _\n" +
		"|_   _|__ _ __ _ __ ___ (_)_ __   __ _| |  _ \\ __ _  __| |\n" +
		"  | |/ _ \\ '__| '_ ` _ \\| | '_ \\ / _` | | |_) / _` |/ _` |\n" +
		"  | |  __/ |  | | | | | | | | | | (_| | |  __/ (_| | (_| |\n" +
		"  |_|\\___|_|  |_| |_| |_|_|_| |_|\\__,_|_|_|   \\__,_|\\__,_|\n"
)

var (
	inputStyle lipgloss.Style = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#aa00ff")).
		Align(lipgloss.Center).
		Width(40).
		Margin(0, 20, 0, 20)
)

type login struct {
	inputs  []textinput.Model
	focused int
	err     error
}

func InitialLogin() *login {
	inputs := make([]textinput.Model, 2)
	inputs[username] = textinput.New()
	inputs[username].Placeholder = "username"
	inputs[username].Focus()
	inputs[username].CharLimit = 15
	inputs[username].Prompt = ""

	inputs[password] = textinput.New()
	inputs[password].Placeholder = "password"
	inputs[password].EchoMode = textinput.EchoPassword
	inputs[password].EchoCharacter = '•'
	inputs[password].CharLimit = 30
	inputs[password].Prompt = ""

	return &login{
		inputs,
		username,
		nil,
	}
}

func (l login) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return textinput.Blink
}

func (l login) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(l.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return l, tea.Quit
		case tea.KeyTab:
			l.focused = (l.focused + 1) % 2
		case tea.KeyShiftTab:
			l.focused = (l.focused - 1) % 2
		}

		for i := range l.inputs {
			l.inputs[i].Blur()
		}
		l.inputs[l.focused].Focus()
	case errMsg:
		l.err = msg
		return l, nil
	}

	for i := range l.inputs {
		l.inputs[i], cmds[i] = l.inputs[i].Update(msg)
	}

	return l, tea.Batch(cmds...)
}

func (l login) View() string {
	s := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fb2f92")).
		AlignHorizontal(lipgloss.Center).
		Width(80).
		Height(8)

	loginStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#d3d3d3")).
		Align(lipgloss.Bottom, lipgloss.Right).
		Width(80).
		Height(4)

	sis := lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Left).
		Width(40).
		Height(2).
		Margin(0, 10, 0, 30)

	tStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#aa00ff"))

	inputsView := ""
	for i, v := range l.inputs {
		var title string
		switch i {
		case username:
			title = "Username"
		case password:
			title = "Password"
		}

		inputsView += sis.Render(
			fmt.Sprintf("%s %s", tStyle.Render(title), v.View()),
		) + "\n"
	}

	return fmt.Sprintf(
		`
%s

%s

%s
        `,
		s.Render(header),
		inputsView,
		loginStyle.Render("Login(Enter)"),
	)
}
