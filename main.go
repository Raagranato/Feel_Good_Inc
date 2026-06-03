package main
//go run main.go
import (
    //"fmt" 
    "FeelGoodInc/internal/ui"
    //"FeelGoodInc/styles"
        
    )

/*TODO:
-Shake tela quando selecionar uma bomba no minesweeper
-Play localy
-Play online(host e client)
-IA do dealer(A cada iteração do turno do dealer, o fluxo de decisão funciona assim:
Calcula a Pontuação: O dealer soma o valor atual de todas as cartas na sua mão.
Avalia a Pontuação:
    Se a pontuação for MENOR que 17: Compra carta obrigatoriamente.
    Se a pontuação for MAIOR que 21: Estourou (Bust). Fim de jogo para ele.
    Se a pontuação for MAIOR que 17 (18, 19, 20, 21): Para (Stand) obrigatoriamente.
    Se a pontuação for EXATAMENTE 17: Aqui entra a checagem:
        É um Hard 17? (Ex: $10 + 7$, ou seja, não tem Ás valendo 11). Ele Para (Stand).
        É um Soft 17?(Ex: $A + 6$, onde o Ás está valendo 11). Ele Compra outra carta (Hit).)
-
-
-
-
-
-
-
-
-
-*/

func main() {
    //fmt.Println(styles.Welcome.Render("Feel Good Inc"))
    ui.Welcome()
    ui.Menu();
}