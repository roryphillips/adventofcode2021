package main

import (
	"adventofcode2021/internal/interpreter"
	"fmt"
	"log"
)

func main() {
	commands, err := interpreter.InterpretInput("./direction_input.txt")
	if err != nil {
		log.Fatal(err)
	}

	hpos, depth, aim := 0, 0, 0
	for _, cmd := range commands {
		switch cmd.Cmd {
		case interpreter.CommandType_Forward:
			hpos += cmd.Val
			depth += aim * cmd.Val
		case interpreter.CommandType_Down:
			aim += cmd.Val
		case interpreter.CommandType_Up:
			aim -= cmd.Val
		default:
			log.Fatalf("unknown command input: %v", err)
		}
	}

	fmt.Println(fmt.Sprintf("Final position is %v", hpos*depth))
}
