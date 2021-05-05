// Créer un fichier executable pour nous quand on va lancer le programme
package main

// La fonction main doit être créer à chaque fois qu'on utilise le package main, la fonctionne sera appelée à chaque fois qu'on execute le programme.
func main() {

	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()
}



