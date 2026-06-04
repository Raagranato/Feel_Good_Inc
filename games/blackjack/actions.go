package blackjack

import (
	"FeelGoodInc/internal/deck"
	"FeelGoodInc/internal/utils"
	"fmt"
	"strings"
)

type Hand struct {
	Cards []deck.Card
	Score int
}

type Player struct { //Vai ter 2 - dealer e player
	Hand
	Name string
	bet int
}

type Dealer struct {
	Hand
}



func (h *Hand) addCard(card deck.Card) {
	h.Cards = append(h.Cards, card)
	h.Score = h.sumPoints()
}

func (h *Hand)isBurst() bool {
	return h.Score > 21
}

func (h *Hand) hasA() bool {
	for i := range h.Cards {
		if "A" == h.Cards[i].Value {
			return true
		}
	}
	return false
}

func (h *Hand) sumPoints() int {
	sum := 0
	contA := 0
	for i := range h.Cards {
		sum += atoi(h.Cards[i].Value)
		if h.Cards[i].Value == "A" {
			contA++
		}
	}
	for sum > 21 && contA > 0 {
		contA--
		sum -= 10
	}

	return sum
}

func atoi(value string) int { //tranforma string do valor da carta para int
	switch value {

	case "A":
		{
			return 11
		}
	case "J", "Q", "K", "10":
		{
			return 10
		}
	case "9":
		{
			return 9
		}
	case "8":
		{
			return 8
		}
	case "7":
		{
			return 7
		}
	case "6":
		{
			return 6
		}
	case "5":
		{
			return 5
		}
	case "4":
		{
			return 4
		}
	case "3":
		{
			return 3
		}
	case "2":
		{
			return 2
		}
	default:
		panic("Erro no atoi(), chegou no final sem identificar uma carta!")
	}

}

//------------------------------------------------------------------------------------------------------------

//Funcoes Blackjack

// func InitPlayers()(player Player,dealer Dealer){

// }

func ReturnBet(player *Player, dealer *Dealer)int{
	switch PlayerWins(player,dealer){
	case 1:{
		utils.YouWon()
		return player.bet*2
	}
	case 0:{
		utils.YouLose()
		return player.bet*-1
	}
	case -1:{
		utils.Draw()
		return player.bet
	}
	}
	panic("Error in function return bet")
}

func PlayerWins(player *Player, dealer *Dealer) int {
    if player.isBurst() {
        return 0
    }
    if dealer.isBurst() {
        return 1
    }
    if player.sumPoints() > dealer.sumPoints() {
        return 1
    }
    if dealer.sumPoints() > player.sumPoints() {
        return 0
    }
    return -1
}

func LoopPlayer(player *Player, dealer *Dealer,mainDeck *[]deck.Card) {
	for {
		if !hit() {
			break
		}
		player.addCard(deck.TakeCard(mainDeck))
		PrintTable(player,dealer)
	}
}

func hit() bool {
	input := string("")
	fmt.Println("Press enter to hit / anything then enter to stay")
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	if input == "" {
		return true // hit
	}
	return false // stand
}

func LoopDealer(dealer *Dealer, mainDeck *[]deck.Card) {
    for hitDealer(dealer) {
        dealer.addCard(deck.TakeCard(mainDeck))
    }
}

func hitDealer(dealer *Dealer) bool {
    return dealer.sumPoints() < 17
}

func PrintCards(cards []deck.Card) {
	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %-2s║ ", card.Value)
	}
	fmt.Println()
	//imprime simbolo
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Suit)
	}
	fmt.Println()
	for i := 0; i < len(cards); i++ {
		fmt.Print("╚═══╝ ")
	}
	fmt.Print("\n\n")
}

func PrintDealer(cards []deck.Card) {

	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Print("╔═══╗ ")
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %-2s║ ", card.Value)
	}
	fmt.Printf("║░░░║ ")
	fmt.Println()
	//imprime simbolo
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Suit)
	}
	fmt.Printf("║░░░║ ")
	fmt.Println()
	for i := 0; i < len(cards); i++ {
		fmt.Print("╚═══╝ ")
	}
	fmt.Print("╚═══╝ ")
	fmt.Print("\n\n")
}
func PrintDealerDone(cards []deck.Card) {
	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %-2s║ ", card.Value)
	}
	fmt.Println()
	//imprime simbolo
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Suit)
	}
	fmt.Println()
	for i := 0; i < len(cards); i++ {
		fmt.Print("╚═══╝ ")
	}
	fmt.Print("\n\n")
}

func PrintTable(player *Player,dealer *Dealer){
	utils.ClearTerminal()
	PrintDealer(dealer.Cards)
	PrintCards(player.Cards)
}

func PrintTableDone(player *Player,dealer *Dealer){
	utils.ClearTerminal()
	PrintDealerDone(dealer.Cards)
	PrintCards(player.Cards)
}