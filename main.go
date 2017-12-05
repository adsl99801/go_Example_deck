package main

import "fmt"

func main() {
	cards := deck{}
	cards = newDeckFromFile("saveResult.txt")
	cards.shuffle()
	fmt.Print(cards.toString())
	fmt.Print("main done")
}
