package main

import (
	"fmt"
)

func main() {
	var choixDuJoueur string
	var cRune []rune
	var bonne int
	mot := Word()
	motDuBot := []rune(mot)
	var tab []int
	var choixJ []rune
	var LetterUse []rune
	tentatives := 10
	fmt.Println("\nBonjour et bienvenu dans notre pendu, le bot a choisi son mot")
	for i := 0; i < len(motDuBot); i++ {
		fmt.Printf("_ ")
	}
	fmt.Printf("\n")
	fmt.Printf("\nTu as %d tentatives pour réussir à trouver le mot que le bot à choisi, bonne chance !\n", tentatives)

	fmt.Println("\nRentre ta lettre :")
	for {
		bonne = 0
		fmt.Println("")
		for {
			fmt.Scan(&choixDuJoueur)
			var diff int
			diff = 0
			choixJ = []rune(choixDuJoueur)
			if len(LetterUse) == 0 {
				break
			}
			if len(choixJ) < 2 {
				for g := 0; g < len(LetterUse); g++ {
					if choixJ[0] == LetterUse[g] {
						fmt.Println("Tu ne peux pas mettre 2 fois la même lettre")
						diff = 1
					}
				}
				if diff == 0 {
					break
				}
			} else {
				fmt.Println("Choisis seulement une lettre")
			}
		}
		cRune = []rune(choixDuJoueur)
		LetterUse = append(LetterUse, choixJ[0])
		for i := 0; i < len(choixDuJoueur); i++ {

			if cRune[i] > 96 && cRune[i] < 123 {
				for i := 0; i < len(motDuBot); i++ {
					if cRune[0] == motDuBot[i] {
						tab = append(tab, i)
						bonne = 1
					}
					if cRune[0] != motDuBot[i] {
					}
				}
			} else {
				fmt.Println("\nTu ne peux choisir seulement des lettres minuscules")
			}
		}
		if bonne == 0 {
			tentatives -= 1
		}
		switch tentatives {
		case 9:
			pendu9()
		case 8:
			pendu8()
		case 7:
			pendu7()
		case 6:
			pendu6()
		case 5:
			pendu5()
		case 4:
			pendu4()
		case 3:
			pendu3()
		case 2:
			pendu2()
		case 1:
			pendu1()
		case 0:
			pendu0()
		}
		fmt.Println("")
		affiche(motDuBot, tab)
		fmt.Printf("\nIl te reste %d tentatives\n\n", tentatives)

		if tentatives > 0 && len(tab) == len(motDuBot) {
			fmt.Println("\nBravo tu as win ")
			break
		}
		if tentatives == 0 {
			fmt.Printf("\nDommage, tu as perdu, le bon mot était : %s", mot)
			break
		}
	}
}

func affiche(bot []rune, tab []int) {

	var pendu []rune
	for i := 0; i < len(bot); i++ {
		pendu = append(pendu, '_')
	}
	for i := 0; i < len(tab); i++ {
		pendu[tab[i]] = bot[tab[i]]
	}
	for i := 0; i < len(pendu); i++ {
		fmt.Printf("%c ", pendu[i])
	}
	fmt.Println("")
}
