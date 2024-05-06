package cmd

import "flag"

type Command struct {
	// Primary name of the command
	Name string
	// Short description of the command
	Synopsis string
	// Included with --help
	HelpText string
	// Supported aliases as alternative entrypoint
	Aliases []string
	// Primary execution of the command
	Run func(c *Command, args []string) error
	// Receives pointer to private Flagset
	// TODO: refactor to account for global, persistent, and local flags
	Flags func(f *flag.FlagSet)
	// Specify expected args and validation
	Arguments Arguments
	// private/internal
	command
}

type command struct {
	// reference to predecessor cmd
	parent *Command
	// descendent cmds that inherit from parent cmd
	subcommands []*Command
	// private flagset, use flagSet() to access
	fs *flag.FlagSet
}

func (c *Command) AddCommand(cmd *Command) {
	// update cmd reference(s)
	cmd.parent = c
	c.subcommands = append(c.subcommands, cmd)
}

func (c *Command) CommandPath() string {
	return c.commandPath()
}

func (c *Command) AppName() string {
	var name string
	for next := c.parent; next != nil; next = next.parent {
		if next.isRootCmd() {
			name = next.Name
		}
	}
	return name
}
