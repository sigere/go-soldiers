package main

import "pbxCodingGojo/army"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func initArmy(a *army.Army, size uint) {
	for _, x := range makeRange(2, 4) {
		for _, y := range makeRange(2, 4) {
			a.Soldiers = append(a.Soldiers, army.Soldier{Compass: army.North, X: x, Y: y})
		}
	}

	a.MapSize = size
}
