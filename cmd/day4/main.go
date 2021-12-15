package main

import (
	"adventofcode2021/internal/systems/bingo"
	"fmt"
	"log"
)

func main() {
	boards, err := bingo.BoardsFromPath("./bingo_input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(boards)
}
