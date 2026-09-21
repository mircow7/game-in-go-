package main

import "game-in-go/structures"

func main() {
	c1 := structures.InitCharacter(
		"Doby",
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion de vie", "Potion de vie", "Potion de vie"},
	)

	c1.DisplayInfo()
}
