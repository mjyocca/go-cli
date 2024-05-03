package cmd

import (
	"fmt"

	"github.com/hashicorp/cli"
)

// wraps cmd.Execute(...args) call
func (cmd *Command) run(args []string) int {
	if cmd.Execute == nil {
		// if no Execute method and has subcommands
		// tell CLI to render helptext output
		if len(cmd.subcommands) != 0 {
			return cli.RunResultHelp
		}
		// otherwise if there are no subcommands return error code
		fmt.Printf("\nError: Command requires registered subcommands or Execute method")
		return 1
	}

	if err := cmd.Execute(cmd, args); err != nil {
		// TODO: handle error
		return 1
	}

	return 0
}
