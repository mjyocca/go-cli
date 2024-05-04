package commands

import (
	"example.com/go-cli/internal/commands/foo"
	"example.com/go-cli/internal/pkg/cmd"
)

// Root Command. Add top level commands with `cmd.AddCommand(...)`
func NewCmdRoot(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "root",
		Synopsis: "<root::synopsis-placeholder>",
		HelpText: "<root::help-text-placeholder>",
	}

	// Add top-level subcommand(s)
	cmd.AddCommand(foo.NewFooCmd(ctx))

	return cmd
}
