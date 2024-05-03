package cmd

import "github.com/hashicorp/cli"

// verify struct satisfies hashicorp/cli interface(s)
var _ cli.Command = (*AdapterCmd)(nil)

// Adapter serving as a mediator between cmd.Command struct{} and hashicorp cli.Command interface{}
type AdapterCmd struct {
	cmd *Command
}

// cli.Command interface{}
func (c *AdapterCmd) Run(args []string) int {
	return c.cmd.run(args)
}

func (c *AdapterCmd) Help() string {
	return c.cmd.HelpText
}

func (c *AdapterCmd) Synopsis() string {
	return c.cmd.Synopsis
}

// Builder method to convert cmd.Command(s) to CommandFactory
func NewCmdFactory(cmd *Command) map[string]cli.CommandFactory {
	m := map[string]cli.CommandFactory{}

	// build alias slice to iterate over
	aliases := []string{cmd.Name}
	for _, alias := range cmd.Aliases {
		aliases = append(aliases, alias)
	}

	// loop over name +aliases
	for _, name := range aliases {
		m[name] = func() (cli.Command, error) {
			return &AdapterCmd{
				cmd: cmd,
			}, nil
		}
	}

	return m
}
