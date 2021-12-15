package navigation

import (
	"adventofcode2021/internal/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func OrdersFromFile(path string) ([]Order, error) {
	orders := []Order{}

	err := input.ReadFileLines(path, func(input string) error {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			return fmt.Errorf("unexpected input: %v", input)
		}
		cmd, val := parts[0], parts[1]

		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert val to int: %v", err)
		}

		order := Order{
			Delta: int(i),
		}
		switch cmd {
		case "forward":
			order.Type = OrderTypeForward
		case "down":
			order.Type = OrderTypeUp
		case "up":
			order.Type = OrderTypeDown
		default:
			return fmt.Errorf("unknown command type: %v", cmd)
		}

		orders = append(orders, order)

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return orders, nil
}
