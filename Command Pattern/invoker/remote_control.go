package invoker

import (
	"command-pattern/command"
	"command-pattern/stack"
)

type RemoteControl struct {
	command command.Command
}

func NewRemoteControl(cmd command.Command) *RemoteControl {
	return &RemoteControl{
		command: cmd,
	}
}

func (rc *RemoteControl) PressButton() {
	stack.Push(rc.command)
	rc.command.PressButton()
}

func (rc *RemoteControl) PressUndo() {
	if !stack.IsEmpty() {
		command, _ := stack.Pop()
		command.PressUndo()
	} else {
		println("No commands to undo")
	}
}
