package main

import (
	"adventofcode2021/internal/systems/sonar"
	"fmt"
	"log"
)

const (
	DefaultWindowSize = 3
)

func main() {
	readings, err := sonar.ReadingsFromFile("./sonar_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	sonarSystem := sonar.NewWindowedSonar(DefaultWindowSize)
	increases := sonarSystem.CalculateDifferences(readings...)

	fmt.Println(fmt.Sprintf("There are %v increases", increases))
}
