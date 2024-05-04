package cmd

import (
	"flag"
	"fmt"

	"github.com/hashicorp/cli"
)

// Invokes Command{} Execute Fn, e.g. `Execute func(c *Command, args []string) error`.
// Responsible for mediating between the two Command{} and cli.Command{} implementations.
// Also handles validations, flag parsing, etc.
func (cmd *Command) run(args []string) int {
	if cmd.Execute == nil {
		// if no Execute method and defined subcommands
		// tell CLI to render helptext output
		if len(cmd.subcommands) != 0 {
			return cli.RunResultHelp
		}
		// otherwise return error code if no subcommands
		fmt.Printf("error: Command requires registered subcommands or Execute method")
		return 1
	}

	// Parse flags
	err := flag.CommandLine.Parse(args)
	if err != nil {
		fmt.Printf("error: problem parsing flags with %q\n", err)
		return 1
	}

	newArgs := flag.CommandLine.Args()

	if err := cmd.Execute(cmd, newArgs); err != nil {
		// TODO: handle error
		return 1
	}

	return 0
}
