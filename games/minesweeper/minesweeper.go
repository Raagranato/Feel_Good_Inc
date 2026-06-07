package minesweeper

//TODO : Fix colors e add color fixar no styles, como amrelo red, blue sla
//colocar cor no quadrado
/*Crash do sistema ao pressionar Enter: Na função PlayerLoop, se o usuário enviar uma string vazia (como apenas apertar a tecla Enter), a função strings.Fields(args) retornará um array vazio. Tentar acessar commands[0] imediatamente na linha seguinte causará um erro fatal de index out of range, derrubando o jogo.

    A mecânica de Bandeira (Flag) apaga as bombas: Quando o jogador digita o comando flag x y, o código define board.boardClosed[x -1][y -1] = -2. Isso sobrescreve o número ou a bomba (o -1) que existia naquela posição. Em um Campo Minado real, a bandeira é apenas uma marcação visual e não deve alterar o conteúdo do tabuleiro.

    Condição de vitória inexistente: A variável board.won é iniciada em 0 e muda para -1 caso você clique em uma bomba. No entanto, não há nenhuma lógica no código que verifique se todas as células seguras foram abertas para alterar o board.won para 1. Por causa disso, o jogador está preso em um loop infinito no minesweeper.go até que ele perca o jogo.

⚠️ Bugs de Lógica do Campo Minado

    Números errados ao redor das bombas: A função CreateBoard() apenas incrementa os valores das células nas direções ortogonais (cima, baixo, esquerda, direita) quando uma bomba é gerada. No Campo Minado, uma bomba afeta todas as 8 direções ao redor dela (incluindo as 4 diagonais).

    Células vazias não abrem as diagonais: A função recursiva openCell também só abre as direções em cruz (x+1, y; x-1, y; x, y+1; x, y-1) quando encontra um 0. Aberturas em cascata também devem atingir as células diagonais para funcionarem corretamente.

    Bandeiras não podem ser removidas: Não existe nenhuma lógica no bloco switch da função PlayerLoop para remover uma bandeira caso o jogador mude de ideia e repita o comando.

    Limites do tabuleiro na marcação de flags: Na função PlayerLoop, o comando flag verifica inBounds(x -1, y-1 ) mas define board.boardClosed[x -1][y -1] = -2 logo em seguida. A função inBounds recebe argumentos inteiros, mas se o usuário digitar letras nas coordenadas, x e y podem assumir valor zero através de erros de conversão não tratados corretamente, fazendo o acesso x-1 e y-1 tentar buscar índices negativos, correndo risco de panic.*/

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

