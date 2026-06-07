package minesweeper



func Play() {
	board := CreateBoard()

	for(board.won == 0){
		printBoard(board)
		PlayerLoop(GetInput(),&board)
	} 
	PrintBoardGameOver(board)
}

