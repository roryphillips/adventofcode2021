package interpreter

import (
	"adventofcode2021/internal/input"
	"fmt"
	"strconv"
	"strings"
)

type CommandType uint32

const (
	CommandType_Unknown CommandType = iota
	CommandType_Forward
	CommandType_Up
	CommandType_Down
)

type Command struct {
	Cmd CommandType
	Val int
}

func InterpretInput(path string) ([]Command, error) {
	commands := []Command{}

	err := input.ReadFileLines(path, func(input string) error {
		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			return fmt.Errorf("unexpected input: %v", input)
		}
		cmd, val := parts[0], parts[1]

		i, err := strconv.ParseInt(val, 10, 32)
		if err != nil {
			return fmt.Errorf("failed to convert val to int: %v", err)
		}

		command := Command{
			Val: int(i),
		}
		switch cmd {
		case "forward":
			command.Cmd = CommandType_Forward
		case "down":
			command.Cmd = CommandType_Down
		case "up":
			command.Cmd = CommandType_Up
		default:
			return fmt.Errorf("unknown command type: %v", cmd)
		}

		commands = append(commands, command)

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	return commands, nil
}
