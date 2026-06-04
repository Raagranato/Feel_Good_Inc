package blackjack


import "FeelGoodInc/internal/deck"

func Play() {
	mainDeck := deck.NewDeck()

	//player := Player{Name: "Jogador 1",}//começa com tudo zerado, n precisa inicializar td

	dealer := Dealer{}

	dealer.addCard(deck.TakeCard(&mainDeck))

	deck.PrintDealer(dealer.Cards)
}
