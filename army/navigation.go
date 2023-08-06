package army

import (
	"errors"
	"fmt"
	"github.com/enescakir/emoji"
	"go-soldiers/vector"
	"strings"
)

type Compass int

type Direction int

type Vector struct {
	x   int
	y   int
	len uint
}

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

func (compass *Compass) String() string {
	switch *compass {
	case North:
		return emoji.UpArrow.String()
	case South:
		return emoji.DownArrow.String()
	case West:
		return emoji.LeftArrow.String()
	case East:
		return emoji.RightArrow.String()
	default:
		panic(fmt.Sprintf("Invalid direction value: %d", *compass))
	}
}

func (compass *Compass) Vector() vector.T {
	switch *compass {
	case North:
		return vector.T{0, 1}
	case South:
		return vector.T{0, -1}
	case West:
		return vector.T{-1, 0}
	case East:
		return vector.T{1, 0}
	default:
		panic(fmt.Sprintf("Invalid compass value: %d", *compass))
	}
}

func (direction *Direction) Vector() vector.T {
	switch *direction {
	case Forward:
		return vector.T{0, 1}
	case Backward:
		return vector.T{0, -1}
	case Left:
		return vector.T{-1, 0}
	case Right:
		return vector.T{1, 0}
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
