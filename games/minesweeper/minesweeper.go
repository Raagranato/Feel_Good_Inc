package minesweeper

import "FeelGoodInc/internal/utils"

//TODO:save?

func Play() {
	board := CreateBoard()

	for board.Won == 0 {
		PrintBoard(board)
		PlayerLoop(GetInput(), &board)
	}

	PrintBoardGameOver(board)
	if board.Won == 1 {
		utils.YouWon()
		// return bet
	}else{
		utils.YouLose()
	}
	
}
