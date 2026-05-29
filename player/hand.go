package players
import "Cards/card"

type Hand struct {
    Cards []card.Card
    Score int
}

type Player struct {
    Hand
    Name string
}

type Dealer struct {
    Hand
}