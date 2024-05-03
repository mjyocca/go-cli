package cmd

type Command struct {
	// Primary name of the command
	Name string
	// Short description of the command
	Synopsis string
	// Included with --help
	HelpText string
	// Supported aliases as alternative entrypoint
	Aliases []string
	// Runs the commands primary logic
	Execute func(c *Command, args []string) error
	// descendent cmds that inherit from root
	subcommands []*Command
}

func (c *Command) AddCommand(cmd *Command) {
	c.subcommands = append(c.subcommands, cmd)
}
