package main

import (
	"github.com/pterm/pterm"
	"go-soldiers/army"
	"go-soldiers/vector"
	"strings"
)

func main() {
	var input string

	pterm.DefaultHeader.WithFullWidth().Println("Hello my commander!")

	commandSet := army.InitCommands()
	var a = army.Army{
		Soldiers:   []army.Soldier{},
		CommandSet: &commandSet,
		Map:        map[vector.T]*army.Soldier{},
		MapSize:    10,
	}
	initArmy(&a)
	var err error
	for {
		_ = display(&a)
		if err != nil {
			pterm.Error.Printfln("Could not execute: %s", err)
		}

		input, _ = pterm.DefaultInteractiveTextInput.Show("On your command")
		split := strings.Split(input, " ")
		err = a.ExecuteCommand(split[0], split[1:])
	}
}
