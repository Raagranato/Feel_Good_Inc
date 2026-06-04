package utils

import ("fmt"
"FeelGoodInc/styles")

func ClearTerminal() {
    fmt.Print("\033[H\033[2J")
}

func YouLose(){
    println(styles.Lose("You Lost! Better luck Next time"))
    println(styles.Lose("(ノಠ益ಠ)ノ彡┻━┻"))
}
func YouWon(){
    println(styles.Win("You won! Nice work"))
    println(styles.Win("(•̀ᴗ•́)و ̑̑"))
}
func Draw(){
    println(styles.Draw("It's a draw"))
    println(styles.Draw("(ಠ_ಠ)"))
}