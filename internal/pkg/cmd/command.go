package cmd

type Command struct {
	Name  string
	Short string
	Long  string

	//
	Execute func(cmd *Command, args []string) error

	Aliases []string
}

func (cmd *Command) run(args []string) int {
	if cmd.Execute == nil {
		return 1
	}

	if err := cmd.Execute(cmd, args); err != nil {
		// handle error
		return 1
	}

	return 0
}
