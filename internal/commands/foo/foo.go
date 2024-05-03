package foo

import "example.com/go-cli/internal/pkg/cmd"

func NewFooCmd(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "foo",
		Synopsis: "<foo::synopsis-placeholder>",
		HelpText: "<foo::help-text-placeholder>",
	}

	cmd.AddCommand(newBarCmd(ctx))
	return cmd
}
