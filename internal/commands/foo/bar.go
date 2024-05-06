package foo

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"strconv"

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
		Synopsis: "All things bar.",
		HelpText: `The "foo bar" command steamlines every and all things bar in relation to foo.`,
		Flags: func(f *flag.FlagSet) {
			f.StringVar(&opts.name, "name", "", "Name of the parameter")
			f.BoolVar(&opts.enabled, "enabled", false, "To enable X setting")
		},
		Arguments: cmd.Arguments{
			Args: []cmd.Argument{
				{
					Name:        "Baz",
					Description: "Name to baz about",
				},
			},
			Validate: cmd.ExactArgs(1),
		},
		Run: func(c *cmd.Command, args []string) error {
			return processInputs(opts, args)
		},
	}

	return cmd
}

// Example accessing flags/arguments
func processInputs(opts *BarOptions, args []string) error {
	resultMap := make(map[string]map[string]string)
	flagMap := map[string]string{
		"name":    opts.name,
		"enabled": strconv.FormatBool(opts.enabled),
	}
	argMap := map[string]string{
		"Baz": args[0],
	}

	resultMap["flags"] = flagMap
	resultMap["arguments"] = argMap

	b, err := json.MarshalIndent(resultMap, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(b))
	return nil
}
