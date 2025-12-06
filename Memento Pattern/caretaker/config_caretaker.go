package caretaker

import (
	"memento-pattern/memento"
	"memento-pattern/originator"
)

type ConfigCaretaker struct {
	mementos []memento.ConfigMemento
}

func NewConfigCaretaker() *ConfigCaretaker {
	return &ConfigCaretaker{}
}

func (cct *ConfigCaretaker) AddMemento(memento memento.ConfigMemento) {
	cct.mementos = append(cct.mementos, memento)
}

func (cct *ConfigCaretaker) Undo(configOriginator *originator.ConfigOriginator) {
	if len(cct.mementos) == 0 {
		return
	}

	last := cct.mementos[len(cct.mementos)-1]

	configOriginator.RestoreMemento(&last)

	cct.mementos = cct.mementos[:len(cct.mementos)-1]
}
