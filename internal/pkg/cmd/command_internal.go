package cmd

import (
	"flag"
	"fmt"
	"slices"
	"strings"

	"github.com/hashicorp/cli"
)

// Invokes Command{} Run Fn, e.g. `Run func(c *Command, args []string) error`.
// Responsible for mediating between the two Command{} and cli.Command{} implementations.
// Also handles validations, flag parsing, etc.
func (cmd *Command) run(args []string) int {
	if cmd.Run == nil {
		// if no Execute method and defined subcommands
		// tell CLI to render helptext output
		if len(cmd.subcommands) != 0 {
			return cli.RunResultHelp
		}
		// otherwise return error code if no subcommands
		fmt.Printf("error: Command requires registered subcommands or Execute method")
		return 1
	}

	// pass flagset ref to cmd if defined
	flagSet := cmd.flagSet()

	// Parse flags
	if err := flagSet.Parse(args); err != nil {
		fmt.Printf("error: problem parsing flagset with %q\n", err)
		return 1
	}

	// Args after flag parsing
	parsedArgs := flagSet.Args()

	// handle Args Validations
	if err := cmd.Arguments.validateFunc()(cmd, parsedArgs); err != nil {
		fmt.Println(err)
		return 1
	}

	// Run the Command
	if err := cmd.Run(cmd, parsedArgs); err != nil {
		// TODO: handle error
		return 1
	}

	return 0
}

// Returns whether command is the root
func (c *Command) isRootCmd() bool {
	return c.parent == nil
}

// Returns full command path to execute the Command.
func (c *Command) commandPath() string {
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

// Returns memoized flagset isolated to this command
func (c *Command) flagSet() *flag.FlagSet {
	if c.fs != nil {
		return c.fs
	}
	flagSet := flag.NewFlagSet(c.commandPath(), flag.ContinueOnError)

	if c.Flags != nil {
		c.Flags(flagSet)
	}

	c.fs = flagSet
	return c.fs
}
