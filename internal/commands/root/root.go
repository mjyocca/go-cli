package root

import (
	"fmt"

	"example.com/go-cli/internal/pkg/cmd"
)

// Root application command
func NewRootCmd(ctx *cmd.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:     "root",
		Synopsis: "<synopsis-placeholder>",
		HelpText: "<help-text-placeholder>",
		Execute: func(c *cmd.Command, args []string) error {
			fmt.Println("root cmd executing with args:", args)
			return nil
		},
	}
	return cmd
}
