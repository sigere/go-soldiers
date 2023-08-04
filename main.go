package main

import (
	"github.com/pterm/pterm"
	"pbxCodingGojo/army"
	"strings"
)

func main() {
	var input string

	pterm.DefaultHeader.WithFullWidth().Println("Hello my commander!")

	commandSet := army.Init()
	var a = army.Army{
		Soldiers:   []army.Soldier{},
		CommandSet: &commandSet,
	}
	initArmy(&a, 10)

	var err error
	for {
		//pterm.Info.Printfln("provided: %s", input)
		_ = display(&a)
		if err != nil {
			pterm.Error.Printfln("Could not execute: %s", err)
		}

		input, _ = pterm.DefaultInteractiveTextInput.Show("On your command")
		split := strings.Split(input, " ")
		err = a.ExecuteCommand(split[0], split[1:])
	}
}
