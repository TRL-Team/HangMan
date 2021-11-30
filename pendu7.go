package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu7() {
	content, err := ioutil.ReadFile("hangman7.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
