package root

import (
	"fmt"

	"example.com/go-cli/internal/pkg/cmd"
)

func NewRootCmd() *cmd.Command {
	cmd := &cmd.Command{
		Name:  "root",
		Short: "<help-short>",
		Long:  "<help-long>",
		Execute: func(c *cmd.Command, args []string) error {
			fmt.Println("hello world!")
			return nil
		},
		// TODO build subcommands
	}

	return cmd
}
