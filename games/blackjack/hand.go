package blackjack
import "FeelGoodInc/internal/deck"

type Hand struct {
    Cards []deck.Card
    Score int
}

type Player struct {
    Hand
    Name string
}

type Dealer struct {
    Hand
}

//func sumPoints(card.cards[]) int num{

//}