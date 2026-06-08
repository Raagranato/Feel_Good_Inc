package minesweeper

//TODO : Fix colors e add color fixar no styles, como amrelo red, blue sla
//colocar cor no quadrado

func Play() {
	board := CreateBoard()

	for(board.won == 0){
		PrintBoard(board)
		PlayerLoop(GetInput(),&board)
	}

	PrintBoardGameOver(board)
	// if (board.won == 1){
	// 	return 30
	// }
}

