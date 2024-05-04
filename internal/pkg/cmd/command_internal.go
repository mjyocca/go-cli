package cmd

import (
	"flag"
	"fmt"
	"slices"
	"strings"

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

	flagSet := flag.NewFlagSet(cmd.pathName(), flag.ContinueOnError)

	// pass flagset ref to cmd if defined
	if cmd.Flags != nil {
		cmd.Flags(flagSet)
	}

	// Parse flags
	if err := flagSet.Parse(args); err != nil {
		fmt.Printf("error: problem parsing flagset with %q\n", err)
		return 1
	}

	// Run
	if err := cmd.Execute(cmd, args); err != nil {
		// TODO: handle error
		return 1
	}

	return 0
}

// If command is the root
func (c *Command) isRootCmd() bool {
	return c.parent == nil
}

// Returns full command path to execute the Command.
func (c *Command) pathName() string {
	cmdNames := []string{c.Name}
	for parent := c.parent; parent != nil; parent = parent.parent {
		// omit for root cmd
		if !parent.isRootCmd() {
			cmdNames = append(cmdNames, parent.Name)
		}
	}
	slices.Reverse(cmdNames)

	return strings.Join(cmdNames, " ")
}
