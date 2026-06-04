package deck

import (
	"fmt"
	"math/rand"
)

func NewDeck() []Card {
	cards := []Card{}
	for i := 0; i < 13; i++ { //13 simbolos
		for j := 0; j < 4; j++ {
			newCard := Card{Suit: Suits[j], Value: Values[i]}
			cards = append(cards, newCard)
		}
	}
	cards = shuffleCards(cards)
	return cards
}

func shuffleCards(cards []Card) []Card {
	for i := len(cards) - 1; i > 0; i-- {
		randomPos := rand.Intn(i + 1)
		cards[i], cards[randomPos] = cards[randomPos], cards[i] //troca 2 sem precisa de um terceiro para armazenar o valor
	}
	return cards
}

func TakeCard(cards *[]Card) Card {
	if len(*cards) == 0 {
		*cards = resetDeck()
	}

	firstCard := (*cards)[0]
	*cards = (*cards)[1:] //pega a primeira o vetor aparti d primeira pos
	return firstCard
}

func resetDeck() []Card {
	return NewDeck()
}

func PrintCards(cards []Card) {
	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Value)
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
	fmt.Print("\n\n\n\n")
}

func PrintDealer(cards []Card) {

	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Print("╔═══╗ ")
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Value)//TODO: o 10 ta quebrando a formatação
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

}
func PrintDealerDone(cards []Card) {
	for i := 0; i < len(cards); i++ {
		fmt.Print("╔═══╗ ")
	}
	fmt.Println()
	//imprime letra
	for _, card := range cards {
		fmt.Printf("║ %s ║ ", card.Value)
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
	fmt.Print("\n\n\n\n")
}
// ╔════════════════════════════════════════╗
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ║                                        ║
// ╚════════════════════════════════════════╝

// │                                        │
// └────────────────────────────────────────┘
