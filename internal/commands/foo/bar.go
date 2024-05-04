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

	cmd := &cmd.Command{
		Name:     "bar",
		Synopsis: "<bar::synopsis-placeholder>",
		HelpText: "<bar::help-text-placeholder>",
		// TODO: make it easier to include flags with HelpText
		Flags: func(f *flag.FlagSet) {
			f.StringVar(&opts.name, "name", "", "Name of the parameter")
			f.BoolVar(&opts.enabled, "enabled", false, "To enable X setting")
		},
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
