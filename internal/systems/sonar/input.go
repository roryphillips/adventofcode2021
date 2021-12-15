package sonar

import (
	"adventofcode2021/internal/utils/input"
	"fmt"
	"strconv"
)

func ReadingsFromFile(path string) ([]int, error) {
	ints := []int{}

	err := input.ReadFileLines(path, func(input string) error {
		i, err := strconv.ParseInt(input, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert to integer: %v", err)
		}

		// Cast is safe here as we're ensuring a 32 bit size
		ints = append(ints, int(i))
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return ints, nil
}
