package main

import (
	"fmt"
	"github.com/enescakir/emoji"
	"go-soldiers/army"
	"go-soldiers/vector"
)

func display(a *army.Army) error {
	board := make([][]string, a.MapSize, a.MapSize)
	for x, _ := range board {
		board[x] = make([]string, a.MapSize)
		for y, _ := range board[x] {

			if soldier, ok := a.Map[vector.T{x, y}]; ok == true {
				board[x][y] = soldier.Compass.String()
			} else {
				board[x][y] = emoji.Herb.String()
			}
		}
	}

	size := int(a.MapSize)
	for y, _ := range board {
		for x, _ := range board[y] {
			_, err := emoji.Printf("%v", board[x][size-y-1])
			if err != nil {
				return err
			}
		}
		fmt.Printf("\n")
	}
	return nil
}
