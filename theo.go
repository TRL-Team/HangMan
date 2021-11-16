package main

import (
	"fmt"
)

func main() {

	var choixDuJoueur string
	var mot string = "cacahuete"
	motDuBot := []rune(mot)
	var tab []int
	//tentatives := 10
	fmt.Println("Rentre ta lettre ou ton mot :")
	fmt.Scan(&choixDuJoueur)
	cRune := []rune(choixDuJoueur)
	for i := 0; i < len(motDuBot); i++ {
		if cRune[0] == motDuBot[i] {
			tab = append(tab, i+1)
		}
		if cRune[0] != motDuBot[i] {
		}
	}
	fmt.Println(tab)
}

// stocker les diffferents choix du joueur dans une variable = motDuJoueur
// for i := 0; motDuJoueur = mot; i ++ {
