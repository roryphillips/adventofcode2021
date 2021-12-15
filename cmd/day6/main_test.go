package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

const TEST_INPUT = "2,5,3,4,4,5,3,2,3,3,2,2,4,2,5,4,1,1,4,4,5,1,2,1,5,2,1,5,1,1,1,2,4,3,3,1,4,2,3,4,5,1,2,5,1,2,2,5,2,4,4,1,4,5,4,2,1,5,5,3,2,1,3,2,1,4,2,5,5,5,2,3,3,5,1,1,5,3,4,2,1,4,4,5,4,5,3,1,4,5,1,5,3,5,4,4,4,1,4,2,2,2,5,4,3,1,4,4,3,4,2,1,1,5,3,3,2,5,3,1,2,2,4,1,4,1,5,1,1,2,5,2,2,5,2,4,4,3,4,1,3,3,5,4,5,4,5,5,5,5,5,4,4,5,3,4,3,3,1,1,5,2,4,5,5,1,5,2,4,5,4,2,4,4,4,2,2,2,2,2,3,5,3,1,1,2,1,1,5,1,4,3,4,2,5,3,4,4,3,5,5,5,4,1,3,4,4,2,2,1,4,1,2,1,2,1,5,5,3,4,1,3,2,1,4,5,1,5,5,1,2,3,4,2,1,4,1,4,2,3,3,2,4,1,4,1,4,4,1,5,3,1,5,2,1,1,2,3,3,2,4,1,2,1,5,1,1,2,1,2,1,2,4,5,3,5,5,1,3,4,1,1,3,3,2,2,4,3,1,1,2,4,1,1,1,5,4,2,4,3"

func benchmarkAges(i int, b *testing.B) {
	chars := strings.Split(TEST_INPUT, ",")
	ages := []int{}
	for _, ch := range chars {
		i, err := strconv.ParseInt(ch, 10, 32)
		if err != nil {
			panic(err)
		}
		ages = append(ages, int(i))
	}

	for n := 0; n < b.N; n++ {
		calculatePopulation(ages, i, 7, 2)
	}
}

func BenchmarkCalculatePopulation(b *testing.B) {
	days := []int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024}
	for _, day := range days {
		b.Run(fmt.Sprintf("Days - %v", day), func(b *testing.B) {
			benchmarkAges(day, b)
		})
	}
}
