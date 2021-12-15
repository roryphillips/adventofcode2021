package diagnostics

import (
	"adventofcode2021/internal/utils/input"
	"fmt"
)

func ReadingsFromFile(path string) ([]string, error) {
	strs := []string{}

	err := input.ReadFileLines(path, func(input string) error {
		strs = append(strs, input)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return strs, nil
}