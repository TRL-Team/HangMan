package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu8() {
	content, err := ioutil.ReadFile("hangman8.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
