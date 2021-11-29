package main

import (
	"fmt"
)

func main() {
	var choixDuJoueur string
	mot := Word()
	motDuBot := []rune(mot)
	var tab []int
	tentatives := 20
	// var appcRune []rune
	fmt.Println("\nBonjour et bienvenu dans notre pendu,")
	for i := 0; i < len(motDuBot); i++ {
		fmt.Printf("_ ")
	}
	fmt.Printf("\n")
	fmt.Println("\nRentre ta lettre ou ton mot :")
	for {
		fmt.Println("")
		fmt.Scan(&choixDuJoueur)
		cRune := []rune(choixDuJoueur)
		//fmt.Printf("%d", cRune)
		// unique(appcRune, choixDuJoueur, cRune)
		for i := 0; i < len(choixDuJoueur); i++ {
			if cRune[i] > 96 && cRune[i] < 123 {
				for i := 0; i < len(motDuBot); i++ {
					if cRune[0] == motDuBot[i] {
						tab = append(tab, i)
						//fmt.Println("GG tu as trouvé une lettre continue !\nRentre une nouvelle lettre :")
					}
					if cRune[0] != motDuBot[i] {
						//fmt.Println("La lettre que tu as choisis n'est pas dans le mot, dommage\nRentre une nouvelle lettre :")
					}
				}
			} else {
				fmt.Println("\nTu ne peux choisir seulement des lettres minuscules")
			}
		}
		affiche(motDuBot, tab)
		tentatives--
		if tentatives <= 0 {
			fmt.Printf("\nDommage, tu as perdu, le bon mot était : %s", mot)
			break
		}
		if len(tab) == len(motDuBot) {
			fmt.Println("bravo vous avez win ")
			break
		}
		tentatives--
		if tentatives <= 0 {
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
		fmt.Printf("%c", pendu[i])
	}
	fmt.Println("")
}
func unique(appcRune []rune, choixDuJoueur string, cRune []rune) {
	// for i := 0; i < len(appcRune); i++ {
	// 	appcRune = append(appcRune, i)
	//}
	for i := 0; i < len(appcRune); i++ {
		if appcRune[i] == cRune[i] {
			appcRune = append(appcRune, cRune...)
			fmt.Println("Tu as dèjà séléctionné cette lettre")
		}
		if appcRune[i] != cRune[i] {
			fmt.Println("non")
		}
	}
}
