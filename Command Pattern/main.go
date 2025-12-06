package main

import (
	"command-pattern/command"
	"command-pattern/invoker"
	"command-pattern/receiver"
	"fmt"
)

func main() {
	onCommand := command.NewOnCommand(*receiver.NewAirConditioner("LG"))
	remoteControl := invoker.NewRemoteControl(onCommand)
	remoteControl.PressButton()

	offCommand := command.NewOffCommand(*receiver.NewAirConditioner("LG"))
	remoteControl = invoker.NewRemoteControl(offCommand)
	remoteControl.PressButton()

	fmt.Print("Undoing last command: ")
	remoteControl.PressUndo()
}
