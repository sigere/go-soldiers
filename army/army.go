package army

import (
	"errors"
	"go-soldiers/vector"
)

type Army struct {
	Soldiers []Soldier
	*CommandSet
	MapSize uint
	Map     map[vector.T]*Soldier
}

func (a *Army) SetSoldiers(soldiers []Soldier) {
	a.Soldiers = soldiers
	for i := range soldiers {
		a.Map[vector.T{soldiers[i].X, soldiers[i].Y}] = &soldiers[i]
	}
}

func (a *Army) ExecuteCommand(name string, args []string) error {
	for _, command := range *a.CommandSet {
		if command.Name == name {
			if err := command.Executable(a, args); err != nil {
				return err
			}
			return nil
		}
	}

	return errors.New("command not found")
}

func (a *Army) TurnToCompass(compass Compass) {
	for i := range a.Soldiers {
		a.Soldiers[i].Turn(compass)
	}
}

func (a *Army) TurnInDirection(direction Direction) error {
	if direction != Right && direction != Left {
		return errors.New("cannot turn in this direction")
	}

	for i := range a.Soldiers {
		a.Soldiers[i].TurnInDirection(direction)
	}

	return nil
}

func (a *Army) Scatter() {
	a.Map = map[vector.T]*Soldier{}
	for i := range a.Soldiers {
		a.Soldiers[i].Scatter()
	}
}

func (a *Army) GoInDirection(direction Direction, steps uint) error {
	newMap := map[vector.T]*Soldier{}
	var err error = nil

	for i := range a.Soldiers {
		soldier := &a.Soldiers[i]
		dir := Direction((int(direction) + int(soldier.Compass)) % 4)
		vec := dir.Vector().Multiply(int(steps))

		_ = soldier.Move(vec, false)
		position := vector.T{soldier.X, soldier.Y}
		_, ok := newMap[vector.T{soldier.X, soldier.Y}]

		if err != nil || ok == true {
			a.restoreSoldiersToMap()
			return errors.New("could not move soldiers")
		}

		newMap[position] = soldier
	}

	a.Map = newMap
	return nil
}

func (a *Army) restoreSoldiersToMap() {
	for position := range a.Map {
		a.Map[position].X = position[0]
		a.Map[position].Y = position[1]
	}
}
