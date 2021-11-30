package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func pendu5() {
	content, err := ioutil.ReadFile("hangman5.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
}
