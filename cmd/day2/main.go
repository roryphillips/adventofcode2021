package main

import (
	"adventofcode2021/internal/systems/navigation"
	"fmt"
	"log"
)

func main() {
	orders, err := navigation.OrdersFromFile("./direction_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	nav := navigation.NewAimNavigation()
	summary, err := nav.ProcessOrders(orders...)
	fmt.Println(summary)
}
