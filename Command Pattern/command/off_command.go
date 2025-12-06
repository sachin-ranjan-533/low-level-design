package command

import "command-pattern/receiver"

type OffCommand struct {
	airConditioner receiver.AirConditioner
}

func NewOffCommand(ac receiver.AirConditioner) *OffCommand {
	return &OffCommand{
		airConditioner: ac,
	}
}

func (oc *OffCommand) PressButton() {
	oc.airConditioner.Off()
}

func (oc *OffCommand) PressUndo() {
	oc.airConditioner.On()
}
