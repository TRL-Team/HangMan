package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu0() {
	content, err := ioutil.ReadFile("hangman0.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
