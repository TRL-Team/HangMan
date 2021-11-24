package main

import (
	"fmt"
)

func main() {
	var pendu []rune
	var choixDuJoueur string
	mot := Word()
	motDuBot := []rune(mot)
	var tab []int
	//tentatives := 10
	fmt.Println("Rentre ta lettre ou ton mot :")
	for i := 0; i < len(motDuBot); i++ {

		for i := 0; i < len(motDuBot); i++ {
			pendu = append(pendu, '_')
		}
		for i := 0; i < len(pendu); i++ {
			fmt.Printf("%c ", pendu[i])
		}
		fmt.Println("")
		fmt.Println(mot)
		fmt.Scan(&choixDuJoueur)
		cRune := []rune(choixDuJoueur)
		for i = 0; i < len(motDuBot); i++ {
			if cRune[0] == motDuBot[i] {
				tab = append(tab, i)
			}
			if cRune[0] != motDuBot[i] {
			}
		}
		pendu[tab[i]] = pendu[cRune[0]]
		for i := 0; i < len(pendu); i++ {
			fmt.Printf("%c ", pendu[i])
		}
	}
}
