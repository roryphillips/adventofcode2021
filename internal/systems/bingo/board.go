package bingo

import "fmt"

func InitBoard(height int, width int) *Board {
	// Width (x) is our first index
	// Height (y) is our second index

	grid := make([][]int, width)
	for i := range grid {
		grid[i] = make([]int, height)
	}

	marks := make([][]int, width)
	for i := range marks {
		marks[i] = make([]int, height)
	}

	return &Board{
		Grid:   grid,
		Marks:  marks,
		Height: height,
		Width:  width,
	}
}

type Board struct {
	Grid   [][]int
	Marks  [][]int
	Height int
	Width  int
	feed   int
}

func (b *Board) LoadRow(cells []int) error {
	if b.feed >= b.Height {
		return fmt.Errorf("attempted to load row beyond range")
	}

	if len(b.Grid[b.feed]) != len(cells) {
		return fmt.Errorf("mismatch in cell lengths, board: %v, input: %v",
			len(b.Grid[b.feed]),
			len(cells))
	}

	for i := range b.Grid[b.feed] {
		b.Grid[b.feed][i] = cells[i]
	}

	b.feed += 1
	return nil
}
