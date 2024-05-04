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
	Flags func(f *flag.FlagSet)
	// reference to predecessor cmd
	parent *Command
	// descendent cmds that inherit from root
	subcommands []*Command
}

func (c *Command) AddCommand(cmd *Command) {
	// update cmd reference(s)
	cmd.parent = c
	c.subcommands = append(c.subcommands, cmd)
}
