package main

import (
	"fmt"
	"github.com/enescakir/emoji"
	"pbxCodingGojo/army"
)

func display(a *army.Army) error {
	board := make([][]string, a.MapSize, a.MapSize)
	for x, _ := range board {
		board[x] = make([]string, a.MapSize)
		for y, _ := range board[x] {
			board[x][y] = emoji.Herb.String()
		}
	}

	for _, soldier := range a.Soldiers {
		board[soldier.X][soldier.Y] = soldier.Compass.String()
	}

	for x, _ := range board {
		for y, _ := range board[x] {
			_, err := emoji.Printf("%v", board[x][y])
			if err != nil {
				return err
			}
		}
		fmt.Printf("\n")
	}
	return nil
}
