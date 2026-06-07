package ui

import (
	//"fmt"
	// "FeelGoodInc/games/blackjack"
	// "FeelGoodInc/games/minesweeper"
	//"FeelGoodInc/internal/utils"
	"FeelGoodInc/styles"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Faz o malabarismo do menu, p ficar bonitinho
type FirstState struct {
	Choices []string
	Opc     int
}

func (m FirstState) Init() tea.Cmd {
	return nil
}

func (m FirstState) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if(m.Opc == 0){
				return m, tea.Quit
			}else if(m.Opc == 1){
				return m, tea.Quit
			}
		case "up":
			if m.Opc != 0 { //impede de passar dos limites
				m.Opc--
			}
		case "down":
			if len(m.Choices) != m.Opc+1 { //gambiarra suprema esse +1
				m.Opc++
			}
		case "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m FirstState) View() string {
	s := styles.Title("Pick your poison") + "\n\n"

	for i := 0; i < len(m.Choices); i++ {
		cursor := "    "
		line := m.Choices[i]
		if i == m.Opc {
			line = styles.Selected(line)
		}
		s += cursor + line + "\n"
	}

	s += "\n(q to exit)"

	return s
}

//-----------------------------------------------------------------------------------------------------

func Menu(){}

func Welcome() {
	fmt.Println(" ╔═══════════════════════════════════════════════════╗ ")
	fmt.Println(" ║░█▀▀░█▀▀░█▀▀░█░░░░░█▀▀░█▀█░█▀█░█▀▄░░░▀█▀░█▀█░█▀▀░░░║ ")
	fmt.Println(" ║░█▀▀░█▀▀░█▀▀░█░░░░░█░█░█░█░█░█░█░█░░░░█░░█░█░█░░░░░║ ")
	fmt.Println(" ║░▀░░░▀▀▀░▀▀▀░▀▀▀░░░▀▀▀░▀▀▀░▀▀▀░▀▀░░░░▀▀▀░▀░▀░▀▀▀░▀░║ ")
	fmt.Println(" ╚═══════════════════════════════════════════════════╝ ")
}
