package main

import (
	"fmt"
)

func main() {

	var joueur int
	fmt.Println("Choisissez une categorie: ")
	fmt.Printf("1)aliment\n2)animaux\n3)marques-voiture\n")
	fmt.Scanln(&joueur)
	switch joueur {
	case 1:
		fmt.Println("aliments")
	case 2:
		fmt.Println("animaux")
	case 3:
		fmt.Println("marques-voiture")
	}

}
