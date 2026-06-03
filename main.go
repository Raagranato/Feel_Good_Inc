package main
//go run main.go
import (
    "fmt" 
    "FeelGoodInc/internal/ui"
    "FeelGoodInc/styles"
        
    )

func main() {
    fmt.Println(styles.Welcome.Render("Feel Good Inc"))
    ui.Menu();
}