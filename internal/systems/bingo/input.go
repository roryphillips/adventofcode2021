package bingo

import (
	"adventofcode2021/internal/utils/input"
	"fmt"
	"strconv"
	"strings"
)

func BoardsFromPath(path string) ([]*Board, error) {
	ticker := ""
	boards := []*Board{}

	var currentBoard *Board
	err := input.ReadFileLines(path, func(input string) error {
		// If we've encountered a new line, initialise our next board
		if input == "" {
			if currentBoard != nil {
				boards = append(boards, currentBoard)
			}
			currentBoard = InitBoard(5, 5)
			return nil
		}

		if strings.Index(input, ",") != -1 {
			ticker = input
			return nil
		}

		cellStrings := strings.Split(input, " ")
		cells := make([]int, len(cellStrings))
		for i := range cells {
			parsed, err := strconv.ParseInt(cellStrings[i], 10, 32)
			if err != nil {
				return fmt.Errorf("failed to parse cell: %v", err)
			}

			cells[i] = int(parsed)
		}

		err := currentBoard.LoadRow(cells)
		if err != nil {
			return fmt.Errorf("failed to load rows into board: %v", err)
		}
		return nil
	})
	// We don't have a final newline so we need to manually append the current board
	boards = append(boards, currentBoard)

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return boards, nil
}
