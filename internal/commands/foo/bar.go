package foo

import (
	"context"
	"flag"
	"fmt"

	"example.com/go-cli/internal/pkg/cmd"
)

// Example local action options
type BarOptions struct {
	ctx     context.Context
	name    string
	enabled bool
}

func newBarCmd(ctx *cmd.Context) *cmd.Command {
	opts := &BarOptions{
		ctx: ctx.ParentCtx,
	}

	// Example using stdlib flag package to parse to this cmd option struct{}
	// TODO: make it easier to include flags with HelpText
	flag.StringVar(&opts.name, "name", "", "")
	flag.BoolVar(&opts.enabled, "enabled", false, "")

	cmd := &cmd.Command{
		Name:     "bar",
		Synopsis: "<bar::synopsis-placeholder>",
		HelpText: "<bar::help-text-placeholder>",
		Execute: func(c *cmd.Command, args []string) error {
			fmt.Println("\nExecute: Bar")
			fmt.Println("--Flags--")
			fmt.Printf("%+v", opts)
			fmt.Println("--Flags--")
			return nil
		},
	}

	return cmd
}
