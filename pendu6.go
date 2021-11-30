package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu6() {
	content, err := ioutil.ReadFile("hangman6.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
