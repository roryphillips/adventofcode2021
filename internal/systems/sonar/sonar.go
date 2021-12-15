package sonar

import "fmt"

type Sonar interface {
	CalculateDifferences(readings ...int) int
}

func NewSimpleSonar() Sonar {
	return &simpleSonar{}
}

type simpleSonar struct {
}

func (s *simpleSonar) CalculateDifferences(readings ...int) int {
	increases := 0
	for i := 1; i < len(readings); i++ {
		a, b := readings[i-1], readings[i]

		if b > a {
			increases += 1
		}
	}

	return increases
}

func NewWindowedSonar(windowSize int) Sonar {
	return &windowedSonar{windowSize: windowSize}
}

type windowedSonar struct {
	windowSize int
}

func (w *windowedSonar) CalculateDifferences(readings ...int) int {
	increases := 0
	// start iteration at the current window size so we can look back one further
	for i := w.windowSize; i < len(readings); i++ {
		windowA := 0
		for offset := w.windowSize; offset > 0; offset-- {
			// if window_size is 3
			// offset = 3
			// offset = 2
			// offset = 1
			windowA += readings[i-offset]
		}
		fmt.Println(windowA)
		windowB := 0
		for offset := w.windowSize - 1; offset >= 0; offset-- {
			// if window_size is 3
			// offset = 2
			// offset = 1
			// offset = 0
			windowB += readings[i-offset]
		}
		fmt.Println(windowB)

		if windowB > windowA {
			increases += 1
		}
	}

	return increases
}
