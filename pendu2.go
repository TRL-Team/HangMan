package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu2() {
	content, err := ioutil.ReadFile("hangman2.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
