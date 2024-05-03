package cmd

import "github.com/hashicorp/cli"

// verify struct satisfies hashicorp/cli interface(s)
var _ cli.Command = (*AdapterCmd)(nil)

// adapter struct to satisfy hashicorp/cli interfaces
type AdapterCmd struct {
	cmd *Command
}

// cli.Command interface{}
func (c *AdapterCmd) Run(args []string) int {
	return c.cmd.Run(args)
}

func (c *AdapterCmd) Help() string {
	return "Help(): <TODO-impl>"
}

func (c *AdapterCmd) Synopsis() string {
	return "Help(): <TODO-impl>"
}
