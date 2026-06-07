package main

//go run main.go
import (
	"FeelGoodInc/games/blackjack"
	"FeelGoodInc/games/minesweeper"
	//"fmt"
	"FeelGoodInc/internal/ui"
	//"FeelGoodInc/styles"
    "FeelGoodInc/internal/utils"
	tea "github.com/charmbracelet/bubbletea"
)

/*TODO:
-Shake tela quando selecionar uma bomba no minesweeper
-Play localy
-Play online(host e client)
-Salvar no meio do jogo
-Upgrades?
-Bet in blackjack(atributo(global?) q é chamado da main e salvo no json, double)
-Jogos:
xadrez
craps
batalha naval
roleta
truco - local deve ser legal
liars dice
horse race
-
-
-
-
-
-*/

func main() {
	//fmt.Println(styles.Welcome.Render("Feel Good Inc"))
	ui.Welcome()

	m := ui.FirstState{
		Choices: []string{"Blackjack", "Minesweeper"},
	}
	result, _ := tea.NewProgram(m).Run()
	finalState := result.(ui.FirstState)
    utils.ClearTerminal()
    utils.SkipLine()
	switch finalState.Opc {
	case 0:
		blackjack.Play()
	case 1:
		minesweeper.Play()
	}
}
