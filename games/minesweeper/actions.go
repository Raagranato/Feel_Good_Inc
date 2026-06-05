package minesweeper

import (
	"FeelGoodInc/styles"
	"fmt"
	"math/rand"
)

const tam = 9

type Board struct {
	boardClosed [tam][tam]int //-1 bomba
	IsOpen      [tam][tam]bool
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
			if inBounds(x+1, y) && board.boardClosed[x+1][y] != -1 {
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
			contBomb++
		}
	}
	return board
}

func inBounds(x, y int) bool {
	return x >= 0 && x < tam && y >= 0 && y < tam
}

func RenderBoard() {

}

func PrintBoardGameOver(board Board) {
    for i := range board.boardClosed {
        for j := range board.boardClosed[i] {
            if board.boardClosed[i][j] == -1 {
                fmt.Print("✸")
            } else {
                fmt.Printf("%d ", board.boardClosed[i][j])
            }
        }
        fmt.Println()
    }
}
func printBoard(board Board) {//TODO:Imprimir os index do lado da matriz
	for i := range board.IsOpen {
		for j := range board.IsOpen[i] {
			if board.IsOpen[i][j] == false {
				fmt.Print("#")
			} else if board.boardClosed[i][j] == -1 {
				fmt.Print(styles.Lose("✸"))
			} else {
				fmt.Printf("%d", board.boardClosed[i][j])
			}
		}
		fmt.Println()
	}
}
