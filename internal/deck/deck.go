package deck
import "math/rand"

func NewDeck()[]Card{
	cards := []Card{}
	for i := 0; i < 13; i++ {//13 simbolos
		for j := 0; j < 4; j++ {
			newCard := Card{Suit: Suits[j], Value: Values[i]}
			cards = append(cards, newCard)
		}
	}
	cards = shuffleCards(cards)
	return cards
}

func shuffleCards(cards []Card)[]Card{
	for i := len(cards) - 1; i > 0; i--{
		randomPos := rand.Intn(i + 1)
		cards[i], cards[randomPos] = cards[randomPos], cards[i]//troca 2 sem precisa de um terceiro para armazenar o valor
	}
	return cards
}

func takeCard(cards *[]Card) Card{
	if len(*cards) == 0 {
        *cards = resetDeck()
    }

	firstCard := (*cards)[0]
	*cards = (*cards)[1:]//pega a primeira o vetor aparti d primeira pos
	return firstCard;
}

func resetDeck() []Card {
    return NewDeck()
}