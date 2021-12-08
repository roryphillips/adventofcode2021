package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type lineFn func(input string) error

func ReadFileLines(path string, onLine lineFn) error {
	file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val := scanner.Text()
		err := onLine(val)
		if err != nil {
			return fmt.Errorf("failed to handle line: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("failed to scan file: %v", err)
	}

	return nil
}

func LoadIntsFromFile(path string) ([]int, error) {
	ints := []int{}

	err := ReadFileLines(path, func(input string) error {
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
