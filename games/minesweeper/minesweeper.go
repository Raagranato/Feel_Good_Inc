package minesweeper

import "FeelGoodInc/internal/utils"

//TODO : Fix colors e add color fixar no styles, como amrelo red, blue sla
//Flag em qualquer coisa
//digitar dois numeros direto xy = 29 x=2 y=9
//timer
//Mostrar quantas flags restam (total de bombas - flags colocadas)
//salvar jogo

func Play() {
	board := CreateBoard()

	for board.won == 0 {
		PrintBoard(board)
		PlayerLoop(GetInput(), &board)
	}

	PrintBoardGameOver(board)
	if board.won == 1 {
		utils.YouWon()
		// return bet
	}else{
		utils.YouLose()
	}
	
}
