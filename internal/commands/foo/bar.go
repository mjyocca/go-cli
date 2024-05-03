package foo

import (
	"fmt"

	"example.com/go-cli/internal/pkg/cmd"
)

func newBarCmd(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "bar",
		Synopsis: "<bar::synopsis-placeholder>",
		HelpText: "<bar::help-text-placeholder>",
		Execute: func(c *cmd.Command, args []string) error {
			fmt.Println("\nbar!!!!")
			return nil
		},
	}

	return cmd
}
