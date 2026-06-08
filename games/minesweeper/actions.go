package minesweeper

import (
	"FeelGoodInc/internal/utils"
	"FeelGoodInc/styles"
	"bufio" //ler linha completa
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

const tam = 9

type Board struct {
	boardClosed [tam][tam]int  //-1 bomba e -2 flag
	IsOpen      [tam][tam]bool //0 fechada 1 open
	won         int //1 se ganhou e -1 se perdeu
}

// número entre 0 e 9
// rand.Intn(10)
func CreateBoard() Board { //inicia as bombas em cada posição
	board := Board{}
	bombMax := (tam * tam) * 15 / 100
	contBomb := 0
	for contBomb < bombMax {
		x := rand.Intn(tam)
		y := rand.Intn(tam)
		if board.boardClosed[x][y] != -1 {
			board.boardClosed[x][y] = -1
			if inBounds(x+1, y) && board.boardClosed[x+1][y] != -1 { //isBomb
				board.boardClosed[x+1][y]++
			}
			if inBounds(x-1, y) && board.boardClosed[x-1][y] != -1 {
				board.boardClosed[x-1][y]++
			}
			if inBounds(x, y+1) && board.boardClosed[x][y+1] != -1 {
				board.boardClosed[x][y+1]++
			}
			if inBounds(x, y-1) && board.boardClosed[x][y-1] != -1 {
				board.boardClosed[x][y-1]++
			}
			if inBounds(x+1, y+1) && board.boardClosed[x+1][y+1] != -1 {
				board.boardClosed[x+1][y+1]++
			}
			if inBounds(x-1, y-1) && board.boardClosed[x-1][y-1] != -1 {
				board.boardClosed[x-1][y-1]++
			}
			if inBounds(x-1, y+1) && board.boardClosed[x-1][y+1] != -1 {
				board.boardClosed[x-1][y+1]++
			}
			if inBounds(x+1, y-1) && board.boardClosed[x+1][y-1] != -1 {
				board.boardClosed[x+1][y-1]++
			}
			contBomb++
		}
	}
	return board
}

func inBounds(x, y int) bool {
	return x >= 0 && x < tam && y >= 0 && y < tam
}

func PlayerLoop(args string, board *Board) {

	commands := strings.Fields(args)
	_, err := strconv.Atoi(commands[0])
	switch {
	case len(commands) == 0:
		{
			fmt.Println("Not a command!")
			fmt.Println("Use: [x] [y] ou flag [x] [y]")
			return
		}
	case commands[0] == "flag":
		{
			if len(commands) != 3 {
				return
			}
			x, err := strconv.Atoi(commands[1])
			y, err := strconv.Atoi(commands[2])
			if err != nil {
				fmt.Println("Coordenada inválida!")
				return
			}
			if inBounds(x-1, y-1) {
				board.boardClosed[x-1][y-1] = -2 //flag spot
			}
		}
	case err == nil:
		{
			x, err := strconv.Atoi(commands[0])
			y, err := strconv.Atoi(commands[1])
			if err != nil {
				fmt.Println("Coordenada inválida!")
				return
			}
			if isBomb(x-1, y-1, board) {
				board.won = -1
			} else {
				openCell(x-1, y-1, board)
			}

		}
	}
}

func openCell(x int, y int, board *Board) {

	if !inBounds(x, y) || board.IsOpen[x][y] || isBomb(x, y, board) {
		return
	}
	board.IsOpen[x][y] = true
	if board.boardClosed[x][y] == 0 {
		openCell(x+1, y, board)
		openCell(x-1, y, board)
		openCell(x, y+1, board)
		openCell(x, y-1, board)
		openCell(x-1, y-1, board)
		openCell(x+1, y-1, board)
		openCell(x+1, y+1, board)
		openCell(x-1, y+1, board)
	}
}

// le linha completa
func GetInput() string {
	fmt.Print("What is your move? ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func isBomb(x int, y int, board *Board) bool {
	if board.boardClosed[x][y] == -1 {
		return true
	}
	return false

}

func PrintBoardGameOver(board Board) {
	fmt.Print("  ")
	for i := range board.IsOpen {
		fmt.Printf("%d ", i+1)
	}
	fmt.Print("\n")
	for i := range board.boardClosed {
		fmt.Printf("%d ", i+1)
		for j := range board.boardClosed[i] {
			if board.boardClosed[i][j] == -1 {
				fmt.Print("✸ ")
			} else {
				fmt.Printf("%d ", board.boardClosed[i][j])
			}
		}
		fmt.Println()
	}
}

func PrintBoard(board Board) {
	utils.ClearTerminal()
	fmt.Print("  ")
	for i := range board.IsOpen {
		fmt.Printf("%d ", i+1)
	}
	fmt.Println()
	for i := range board.IsOpen {
		fmt.Printf("%d ", i+1)
		for j := range board.IsOpen[i] {

			if board.boardClosed[i][j] == -2 {
				fmt.Print(styles.Flag("⚑ "))
			} else if board.IsOpen[i][j] == false {
				fmt.Print(styles.Closed("▣ "))
			} else if board.boardClosed[i][j] == -1 {
				fmt.Print(styles.Lose("✸ "))
			} else {
				switch board.boardClosed[i][j] {
				case 0:
					fmt.Print("  ")
				case 1:
					fmt.Print(styles.One("1 "))
				case 2:
					fmt.Print(styles.Two("2 "))
				case 3:
					fmt.Print(styles.Three("3 "))
				case 4:
					fmt.Print(styles.Four("4 "))
				case 5:
					fmt.Print(styles.Five("5 "))
				case 6:
					fmt.Print(styles.Six("6 "))
				case 7:
					fmt.Print(styles.Seven("7 "))
				case 8:
					fmt.Print(styles.Eight("8 "))
				}
			}
		}
		fmt.Println()
	}
}
