//model do bubble tea é uma interface que exige init update view

package gui

import (
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func NewProgram() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	choices   []string
	options   []string
	cursor    int
	userInput textinput.Model
}

func initialModel() model {

	i := textinput.New()
	i.Prompt = ""
	i.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	i.Width = 48
	i.SetValue("")
	i.CursorEnd()

	return model{
		choices: []string{"Novo", "Rodar", "Deletar"},

		userInput: i,

		options: []string{},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c":
			return m, tea.Quit

		case "left":
			if m.cursor >= 1 {
				m.cursor--
			}

		case "right":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			switch m.choices[m.cursor] {

			case "Novo":
				m.userInput.Focus()

				var cmd tea.Cmd
				m.userInput, cmd = m.userInput.Update(msg)
				return m, cmd

			case "Rodar":
				fmt.Print("Rodar")
			case "Deletar":
				fmt.Print("deletar")

			}
		}
	}
	return m, nil
}

func (m model) View() string {
	s := Title() + "\n\n"

	s += Options(m.options)

	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = Cursor()
		}

		s += fmt.Sprintf("%s |%s|", cursor, choice)

	}

	s += "\n" + m.userInput.View() + "\n\n"

	s += "\n\nAperte CTRL+C para sair\n"

	return s
}
