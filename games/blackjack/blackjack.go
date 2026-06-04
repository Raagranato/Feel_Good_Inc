package blackjack

import (
	"FeelGoodInc/internal/deck"
	//"FeelGoodInc/internal"
	//"fmt"
	
)

func Play() {
	mainDeck := deck.NewDeck()

	player := Player{Name: "Jogador 1"} //começa com tudo zerado, n precisa inicializar td

	dealer := Dealer{}

	dealer.addCard(deck.TakeCard(&mainDeck))
	player.addCard(deck.TakeCard(&mainDeck))
	player.addCard(deck.TakeCard(&mainDeck))
	
	PrintDealer(dealer.Cards)
	PrintCards(player.Cards)
	
	LoopPlayer(&player,&dealer,&mainDeck)
	
	LoopDealer(&dealer,&mainDeck)
	PrintTableDone(&player,&dealer)
	ReturnBet(&player,&dealer)

}
