package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu3() {
	content, err := ioutil.ReadFile("hangman3.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
