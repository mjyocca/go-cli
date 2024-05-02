package cmd

import "github.com/hashicorp/cli"

// verify struct satisfies hashicorp/cli interface(s)
var _ cli.Command = (*commandCli)(nil)

type commandCli struct {
	cmd *Command
}

// cli.Command interface{}
func (c *commandCli) Run(args []string) int {
	return c.cmd.run(args)
}

func (c *commandCli) Help() string {
	return "Help(): <TODO-impl>"
}

func (c *commandCli) Synopsis() string {
	return "Help(): <TODO-impl>"
}
