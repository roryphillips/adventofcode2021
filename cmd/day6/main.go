package main

import (
	"adventofcode2021/internal/problems/laternfish"
	"fmt"
	"log"
)

func main() {
	ages, err := laternfish.AgesFromPath("./lantern_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(calculatePopulation(ages, 512, 7, 2))
}
