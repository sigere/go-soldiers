package army

import (
	"errors"
	"math/rand"
)

type Executable func(army *Army, args []string) error

func Turn(army *Army, args []string) error {
	if len(args) < 1 {
		return errors.New("this command requires one argument")
	}

	compass, err := CompassFromString(args[0])
	if err != nil {
		return err
	}

	for i := range army.Soldiers {
		army.Soldiers[i].Compass = compass
	}

	return nil
}

func Scatter(army *Army, arg []string) error {
	for i := range army.Soldiers {
		army.Soldiers[i].Compass = Compass(rand.Intn(4))
		army.Soldiers[i].X = rand.Intn(int(army.MapSize))
		army.Soldiers[i].Y = rand.Intn(int(army.MapSize))
	}

	return nil
}

func (a *Army) Forward() error {
	for i := range a.Soldiers {
		switch a.Soldiers[i].Compass {
		case North:
			a.Soldiers[i].X--
		case East:
			a.Soldiers[i].Y++
		case South:
			a.Soldiers[i].X++
		case West:
			a.Soldiers[i].Y--
		}
	}
	return nil
}
