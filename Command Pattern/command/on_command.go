package command

import "command-pattern/receiver"

type OnCommand struct {
	airConditioner receiver.AirConditioner
}

func NewOnCommand(ac receiver.AirConditioner) *OnCommand {
	return &OnCommand{
		airConditioner: ac,
	}
}

func (oc *OnCommand) PressButton() {
	oc.airConditioner.On()
}

func (oc *OnCommand) PressUndo() {
	oc.airConditioner.Off()
}
