package main

import (
	"github.com/stretchr/testify/assert"
	"go-soldiers/army"
	"testing"
)

func initTestArmy() army.Army {
	commandSet := army.InitCommands()
	return army.Army{
		Soldiers: []army.Soldier{
			{
				Compass: army.North,
				X:       1,
				Y:       1,
			},
			{
				Compass: army.West,
				X:       2,
				Y:       2,
			},
			{
				Compass: army.East,
				X:       3,
				Y:       3,
			},
		},
		CommandSet: &commandSet,
		MapSize:    10,
	}
}

func Test(t *testing.T) {
	a := initTestArmy()
	dataSet := []struct {
		compass army.Compass
		command []string
	}{
		{army.East, []string{"turn", "east"}},
		{army.West, []string{"turn", "west"}},
		{army.South, []string{"turn", "south"}},
		{army.North, []string{"turn", "north"}},
	}

	for _, value := range dataSet {
		_ = a.ExecuteCommand(value.command[0], value.command[1:])

		for i := range a.Soldiers {
			assert.Equal(t, value.compass, a.Soldiers[i].Compass, "Soldier was not turned south")
		}
	}
}
