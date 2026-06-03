package ui

import (
	//"fmt"
	"FeelGoodInc/internal/utils"
	"FeelGoodInc/styles"
	tea "github.com/charmbracelet/bubbletea"
	"fmt"
)

// Faz o malabarismo do menu, p ficar bonitinho
type FirstState struct {
	choices []string
	opc     int
}

func (m FirstState) Init() tea.Cmd {
	return nil
}

func (m FirstState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.opc != 0 { //impede de passar dos limites
				m.opc--
			}
		case "down":
			if len(m.choices) != m.opc+1 { //gambiarra suprema esse +1
				m.opc++
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m FirstState) View() string {
	s := styles.Title.Render("Pick your poison") + "\n\n"

	for i := 0; i < len(m.choices); i++ {
		cursor := "    "
		line := m.choices[i]
		if i == m.opc {
			line = styles.Selected.Render(line)
		}
		s += cursor + line + "\n"
	}

	s += "\n(q to exit)"

	return s
}

//-----------------------------------------------------------------------------------------------------

func Menu() {
	m := FirstState{
		choices: []string{"Blackjack", "Mineswipe"},
		opc:     0,
	}
	tea.NewProgram(m).Run()
	utils.ClearTerminal()
}

func Welcome() {
	fmt.Println(" ╔═══════════════════════════════════════════════════╗ ")
	fmt.Println(" ║░█▀▀░█▀▀░█▀▀░█░░░░░█▀▀░█▀█░█▀█░█▀▄░░░▀█▀░█▀█░█▀▀░░░║ ")
	fmt.Println(" ║░█▀▀░█▀▀░█▀▀░█░░░░░█░█░█░█░█░█░█░█░░░░█░░█░█░█░░░░░║ ")
	fmt.Println(" ║░▀░░░▀▀▀░▀▀▀░▀▀▀░░░▀▀▀░▀▀▀░▀▀▀░▀▀░░░░▀▀▀░▀░▀░▀▀▀░▀░║ ")
	fmt.Println(" ╚═══════════════════════════════════════════════════╝ ")
}
