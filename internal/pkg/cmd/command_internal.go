package cmd

// wraps cmd.Execute(...args) call
func (cmd *Command) run(args []string) int {
	if cmd.Execute == nil {
		return 1
	}

	if err := cmd.Execute(cmd, args); err != nil {
		// TODO: handle error
		return 1
	}

	return 0
}
