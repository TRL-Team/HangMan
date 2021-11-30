package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu9() {
	content, err := ioutil.ReadFile("hangman9.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
