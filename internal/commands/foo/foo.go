package foo

import "example.com/go-cli/internal/pkg/cmd"

func NewFooCmd(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "foo",
		Synopsis: "All things Foo",
		HelpText: `Command group streamlining all things foo`,
	}

	cmd.AddCommand(newBarCmd(ctx))
	return cmd
}
