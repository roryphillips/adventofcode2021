package main

import (
	"adventofcode2021/internal/input"
	"fmt"
	"log"
)

const (
	DefaultWindowSize = 3
)

func main() {
	ints, err := input.LoadIntsFromFile("./sonar_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	increases := calculateIncreasesWindowed(DefaultWindowSize, ints)

	fmt.Println(fmt.Sprintf("There are %v increases", increases))
}

func calculateIncreasesSimple(ints []int) int {
	increases := 0
	for i := 1; i < len(ints); i++ {
		a, b := ints[i-1], ints[i]

		if b > a {
			increases += 1
		}
	}

	return increases
}

func calculateIncreasesWindowed(windowSize int, ints []int) int {
	increases := 0
	// start iteration at the current window size so we can look back one further
	for i := windowSize; i < len(ints); i++ {
		windowA := 0
		for offset := windowSize; offset > 0; offset-- {
			// if window_size is 3
			// offset = 3
			// offset = 2
			// offset = 1
			windowA += ints[i-offset]
		}
		fmt.Println(windowA)
		windowB := 0
		for offset := windowSize - 1; offset >= 0; offset-- {
			// if window_size is 3
			// offset = 2
			// offset = 1
			// offset = 0
			windowB += ints[i-offset]
		}
		fmt.Println(windowB)

		if windowB > windowA {
			increases += 1
		}
	}

	return increases
}
