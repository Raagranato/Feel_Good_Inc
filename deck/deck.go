package deck
import "Cards/card"
import "math/rand"

func newDeck()[]card.Card{
	cards := []card.Card{}
	for i := 0; i < 13; i++ {//13 simbolos
		for j := 0; j < 4; j++ {
			newCard := card.Card{Suit: card.Suits[j], Value: card.Values[i]}
			cards = append(cards, newCard)
		}
	}
	return cards
}

func shufleCards(cards []card.Card)[]card.Card{
	for i := len(cards) - 1; i > 0; i--{
		randomPos := rand.Intn(i + 1)
		cards[i], cards[randomPos] = cards[randomPos], cards[i]//troca 2 sem precisa de um terceiro para armazenar o valor
	}
	return cards
}