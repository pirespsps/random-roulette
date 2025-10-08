package gui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func Menu(op []string) {

	line()

	Options(op)

	instructions()

}

func Resultado(result string) {

	border := lipgloss.Border{
		Top:         "_",
		Bottom:      "",
		Left:        "|",
		Right:       "|",
		BottomLeft:  "/",
		BottomRight: "\\",
		TopLeft:     ".",
		TopRight:    ".",
	}

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Padding(2).
		Background(lipgloss.Color("#7D56F4")).
		Border(border).
		Height(2)

	fmt.Println(style.Render(result))
}

func Options(options []string) string {

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#e1027e"))

	text := "Valores: "

	for _, v := range options {

		text += style.Render(v)
		text += lipgloss.NewStyle().Foreground(lipgloss.Color("#7D56F4")).Render(" | ")
	}

	return lipgloss.NewStyle().Render(text + "\n\n")
}

func instructions() {
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4"))

	barStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#e1027e"))

	text := "Digite o valor"
	text += barStyle.Render(" | ")
	text += style.Render("spin")
	text += " para rodar"
	text += barStyle.Render(" | ")
	text += style.Render("q")
	text += " para sair"

	fmt.Println(lipgloss.NewStyle().MarginBottom(1).Render(text))
}

func Title() string {

	var myCuteBorder = lipgloss.Border{
		Top:         "._.:*:",
		Bottom:      "._.:*:",
		Left:        "|*",
		Right:       "|*",
		TopLeft:     "*",
		TopRight:    "*",
		BottomLeft:  "*",
		BottomRight: "*",
	}

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		PaddingTop(1).
		PaddingBottom(1).
		PaddingLeft(5).
		Width(22).
		Border(myCuteBorder).
		BorderForeground(lipgloss.Color("#e1027e"))

	return style.Render("Roulette\n")
}

func line() {
	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#e1027e"))

	fmt.Println(style.Render("\n-----------------------------------------------\n"))
}
