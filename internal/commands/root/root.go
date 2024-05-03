package root

import (
	"example.com/go-cli/internal/commands/foo"
	"example.com/go-cli/internal/pkg/cmd"
)

// Root Command. Add top level commands with `cmd.AddCommand(...)`
func NewRootCmd(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "root",
		Synopsis: "<root::synopsis-placeholder>",
		HelpText: "<root::help-text-placeholder>",
	}

	cmd.AddCommand(foo.NewFooCmd(ctx))

	return cmd
}
