package main

import (
	"adventofcode2021/internal/systems/diagnostics"
	"fmt"
	"log"
)

func main() {
	readings, err := diagnostics.ReadingsFromFile("./diagnostic_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	diag := diagnostics.NewDiagnostics(12)
	powerUsage, err := diag.CalculatePowerUsage(readings...)
	if err != nil {
		log.Fatal(err)
	}
	lifeSupport, err := diag.CalculateLifeSupportRating(readings...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(powerUsage)
	fmt.Println(lifeSupport)
}
