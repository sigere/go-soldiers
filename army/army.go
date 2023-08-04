package army

import (
	"errors"
)

type Army struct {
	Soldiers []Soldier
	*CommandSet
	MapSize uint
}

func (army *Army) ExecuteCommand(name string, args []string) error {
	for _, command := range *army.CommandSet {
		if command.Name == name {
			if err := command.Executable(army, args); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("command not found")
}
