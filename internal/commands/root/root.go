package root

import (
	"context"
	"fmt"

	"example.com/go-cli/internal/pkg/cmd"
)

func NewRootCmd(ctx context.Context) *cmd.Command {
	cmd := &cmd.Command{
		Name:  "root",
		Short: "<help-short>",
		Long:  "<help-long>",
		Execute: func(c *cmd.Command, args []string) error {
			fmt.Println("Cmd.Execute(root)", args)
			return nil
		},
	}
	return cmd
}
