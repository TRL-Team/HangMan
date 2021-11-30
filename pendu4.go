package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu4() {
	content, err := ioutil.ReadFile("hangman4.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
