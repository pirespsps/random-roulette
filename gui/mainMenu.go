//model do bubble tea é uma interface que exige init update view

package gui

import (
	"fmt"

	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/lipgloss"

	tea "github.com/charmbracelet/bubbletea"
)

func NewProgram() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	choices  []string
	options  []string
	textarea textarea.Model
	cursor   int
	selected map[int]struct{}
}

func initialModel() model {
	ta := textarea.New()
	ta.Placeholder = "Adicione um item..."
	ta.Focus()
	ta.Prompt = "|"
	ta.CharLimit = 280
	ta.SetWidth(30)
	ta.SetHeight(1)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	return model{
		choices: []string{"Novo", "Rodar", "Deletar"},

		textarea: ta,

		options: []string{},

		selected: make(map[int]struct{}),
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "left", "a":
			if m.cursor >= 1 {
				m.cursor--
			}

		case "right", "d":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			switch m.selected[m.cursor] {
			//
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
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s ", cursor, checked, choice)
	}

	s += "\n\n" + m.textarea.View()

	s += "\n\nAperte Q para sair\n"

	return s
}
