package input

import (
	"bufio"
	"fmt"
	"os"
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
