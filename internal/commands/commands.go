package commands

import (
	"example.com/go-cli/internal/commands/foo"
	"example.com/go-cli/internal/commands/qux"
	"example.com/go-cli/internal/pkg/cmd"
)

// Root Command. Add top level commands with `cmd.AddCommand(...)`
func NewCmdRoot(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     ctx.AppName,
		HelpText: `CLI Application written in Go/Golang`,
	}

	// Add top-level subcommand(s)
	cmd.AddCommand(foo.NewFooCmd(ctx))
	cmd.AddCommand(qux.NewCmdQux(ctx))

	return cmd
}
