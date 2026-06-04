package blackjack

import (
	"FeelGoodInc/internal/deck"
)

type Hand struct {
	Cards []deck.Card
	Score int
}

type Player struct { //Vai ter 2 - dealer e player
	Hand
	Name string
}

type Dealer struct{
    Hand
}

func (h *Hand) addCard(card deck.Card) {
	h.Cards = append(h.Cards, card)
	h.Score = h.sumPoints()
}

func isBurst(h *Hand) bool {
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
        if h.Cards[i].Value == "A"{
            contA++
        }
	}
    for sum > 21 && contA > 0{
        contA--
        sum-=10
    }
    
	return sum
}

func atoi(value string) int { //tranforma string do valor da carta para int
	switch value {

	case "A":
		{
			return 11
		}
	case "J", "Q", "K","10":
		{
			return 10
		}
	case "9":
		{
			return  9
		}
	case "8":
		{
			return  8
		}
	case "7":
		{
			return  7
		}
	case "6":
		{
			return  6
		}
	case "5":
		{
			return  5
		}
	case "4":
		{
			return  4
		}
	case "3":
		{
			return  3
		}
	case "2":
		{
			return  2
		}
	default:
		panic("Erro no atoi(), chegou no final sem identificar uma carta!")
	}

}
//------------------------------------------------------------------------------------------------------------



// func InitPlayers()(player Player,dealer Dealer){
    
// }