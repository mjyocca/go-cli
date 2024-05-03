package cmd

import (
	"fmt"

	"github.com/hashicorp/cli"
)

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

// Builder method to convert cmd.Command(s) to CommandFactory map
// Accepts root Command{} and traverses subcommands to fill CommandFactory
func NewCmdFactory(cmd *Command) map[string]cli.CommandFactory {
	factory := map[string]cli.CommandFactory{}

	for _, subCmd := range cmd.subcommands {
		traverseSubCmds("", subCmd, factory)
	}

	return factory
}

// Recursively traverses cmd subcommand tree to fill the command factory map along with possible aliasess
func traverseSubCmds(parentPath string, cmd *Command, factory map[string]cli.CommandFactory) {
	// build alias slice to iterate over
	aliases := []string{cmd.Name}
	for _, alias := range cmd.Aliases {
		aliases = append(aliases, alias)
	}

	// loop over name +aliases to fill factory map
	for _, name := range aliases {
		cmdPath := resolveCmdPath(parentPath, name)

		factory[cmdPath] = buildCmdFactory(cmd)

		for _, sub := range cmd.subcommands {
			traverseSubCmds(cmdPath, sub, factory)
		}
	}
}

// Utility to include parent or super command
func resolveCmdPath(parent string, current string) string {
	if parent == "" {
		return current
	}
	return fmt.Sprintf("%s %s", parent, current)
}

// Wraps Command{} Adapter implementation
func buildCmdFactory(cmd *Command) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return &AdapterCmd{
			cmd: cmd,
		}, nil
	}
}
