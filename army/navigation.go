package army

import (
	"errors"
	"fmt"
	"github.com/enescakir/emoji"
	"strings"
)

type Compass int

type Direction int

const (
	North Compass = iota
	East
	South
	West
)

const (
	Forward Direction = iota
	Right
	Backward
	Left
)

func (direction *Compass) String() string {
	switch *direction {
	case North:
		return emoji.UpArrow.String()
	case South:
		return emoji.DownArrow.String()
	case West:
		return emoji.LeftArrow.String()
	case East:
		return emoji.RightArrow.String()
	default:
		panic(fmt.Sprintf("Invalid direction value: %d", *direction))
	}
}

func CompassFromString(str string) (Compass, error) {
	switch strings.ToLower(str) {
	case "north":
		return North, nil
	case "east":
		return East, nil
	case "south":
		return South, nil
	case "west":
		return West, nil
	}

	return North, errors.New(fmt.Sprintf(
		"%s is not valid geographical direction", str))
}

func DirectionFromString(str string) (Direction, error) {
	switch strings.ToLower(str) {
	case "forward":
		return Forward, nil
	case "right":
		return Right, nil
	case "backwards":
		return Backward, nil
	case "left":
		return Left, nil
	}

	return Forward, errors.New(fmt.Sprintf(
		"%s is not valid direction", str))
}
