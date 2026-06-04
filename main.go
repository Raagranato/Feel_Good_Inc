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
-Salvar no meio do jogo
-Upgrades?
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