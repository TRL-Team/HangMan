package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu1() {
	content, err := ioutil.ReadFile("hangman1.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
