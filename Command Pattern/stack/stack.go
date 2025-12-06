package stack

import "command-pattern/command"

var stack = make([]command.Command, 0)

func Push(command command.Command) {
	stack = append(stack, command)
}

func Pop() (command.Command, bool) {
	if len(stack) == 0 {
		return nil, false
	}
	lastIndex := len(stack) - 1
	cmd := stack[lastIndex]
	stack = stack[:lastIndex]
	return cmd, true
}

func IsEmpty() bool {
	return len(stack) == 0
}
