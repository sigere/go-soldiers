package main

import "go-soldiers/army"

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func initArmy(a *army.Army) {
	
	var soldiers []army.Soldier
	for _, x := range makeRange(2, 4) {
		for _, y := range makeRange(7, 9) {
			soldiers = append(soldiers, army.Soldier{
				Compass: army.North, X: x, Y: y, Army: a,
			})
		}
	}
	a.SetSoldiers(soldiers)
}
