package army

import (
	"math/rand"
	"time"
)

var generator rand.Rand

type Command struct {
	Name       string
	Executable Executable
}

type CommandSet map[string]Command

func InitCommands() CommandSet {
	return CommandSet{
		"turn": {
			Name:       "turn",
			Executable: TurnCommand,
		},
		"scatter": {
			Name:       "scatter",
			Executable: ScatterCommand,
		},
		"go": {
			Name:       "go",
			Executable: GoInDirectionCommand,
		},
	}
}

func init() {
	generator = *rand.New(rand.NewSource(time.Now().UnixNano()))
}
