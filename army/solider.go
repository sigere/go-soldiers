package army

import (
	"errors"
	"fmt"
	"go-soldiers/vector"
)

type Soldier struct {
	Compass Compass
	X       int
	Y       int
	Army    *Army
}

func (s *Soldier) Turn(compass Compass) {
	s.Compass = compass
}

func (s *Soldier) TurnInDirection(direction Direction) {
	var addition uint
	if direction == Right {
		addition = 1
	} else if direction == Left {
		addition = 3
	}

	s.Compass = Compass((uint(s.Compass) + addition) % 4)
}

func (s *Soldier) Scatter() {
	var position vector.T
	for taken := true; taken == true; { // look for free place for solider
		s.Compass = Compass(generator.Intn(4))
		s.X = generator.Intn(int(s.Army.MapSize))
		s.Y = generator.Intn(int(s.Army.MapSize))

		position = vector.T{s.X, s.Y}
		_, taken = s.Army.Map[position]
	}
	s.Army.Map[position] = s
}

func (s *Soldier) Move(vec vector.T, updateMap bool) error {
	newX, newY := s.X+vec[0], s.Y+vec[1]
	mapSize := s.Army.MapSize
	if newX < 0 || uint(newX) >= mapSize || newY < 0 || uint(newY) >= mapSize {
		return errors.New(fmt.Sprintf("position (%d, %d) is out of scope", newX, newY))
	}

	if updateMap {
		delete(s.Army.Map, vector.T{s.X, s.Y})
		s.Army.Map[vector.T{newX, newY}] = s
	}

	s.X = newX
	s.Y = newY

	return nil
}
