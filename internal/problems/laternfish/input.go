package laternfish

import (
	"adventofcode2021/internal/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func AgesFromPath(path string) ([]int, error) {
	ints := []int{}

	err := input.ReadFileLines(path, func(input string) error {
		split := strings.Split(input, ",")
		for _, i := range split {
			num, err := strconv.ParseInt(i, 10, 32)
			if err != nil {
				return fmt.Errorf("failed to convert to integer: %v", err)
			}

			//Cast is safe here as we're ensuring a 32 bit size
			ints = append(ints, int(num))
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return ints, nil
}
