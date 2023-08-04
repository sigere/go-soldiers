package army

type Command struct {
	Name       string
	Executable Executable
}

type CommandSet []Command

func Init() CommandSet {
	return CommandSet{
		{
			Name:       "turn",
			Executable: Turn,
		},
		{
			Name:       "scatter",
			Executable: Scatter,
		},
	}
}
