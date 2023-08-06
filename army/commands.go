package army

import (
	"errors"
	"strconv"
)

type Executable func(a *Army, args []string) error

func TurnCommand(a *Army, args []string) error {
	if len(args) < 1 {
		return errors.New("this command requires one argument")
	}

	compass, err := CompassFromString(args[0])
	if err == nil {
		a.TurnToCompass(compass)
		return nil
	}

	direction, err := DirectionFromString(args[0])
	if err == nil {
		return a.TurnInDirection(direction)
	}

	return err
}

func ScatterCommand(a *Army, args []string) error {
	a.Scatter()
	return nil
}

func GoInDirectionCommand(a *Army, args []string) error {
	if len(args) < 1 {
		return errors.New("this command requires at least one argument")
	}

	direction, err := DirectionFromString(args[0])
	if err != nil {
		return err
	}

	var steps uint = 1
	if len(args) >= 2 {
		converted, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}
		if converted < 0 {
			return errors.New("number of steps must be positive")
		}

		steps = uint(converted)
	}

	return a.GoInDirection(direction, steps)
}
