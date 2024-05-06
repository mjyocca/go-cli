package cmd

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/hashicorp/cli"
)

// Builder method to convert cmd.Command(s) to CommandFactory map
// Accepts root Command{} and traverses subcommands to fill CommandFactory
func NewCmdFactory(cmd *Command) map[string]cli.CommandFactory {
	factory := map[string]cli.CommandFactory{}

	for _, subCmd := range cmd.subcommands {
		traverseSubCmds("", subCmd, factory)
	}

	return factory
}

// Recursively traverses cmd subcommand tree to fill the command factory map along with possible aliasess
func traverseSubCmds(parentPath string, cmd *Command, factory map[string]cli.CommandFactory) {
	// build alias slice to iterate over
	aliases := []string{cmd.Name}
	for _, alias := range cmd.Aliases {
		aliases = append(aliases, alias)
	}

	// loop over name + aliases to fill factory map
	for _, name := range aliases {
		cmdPath := resolveCmdPath(parentPath, name)

		factory[cmdPath] = buildCmdFactory(cmd)

		for _, sub := range cmd.subcommands {
			traverseSubCmds(cmdPath, sub, factory)
		}
	}
}

// Utility to include parent or super command
func resolveCmdPath(prev string, current string) string {
	if prev == "" {
		return current
	}
	return fmt.Sprintf("%s %s", prev, current)
}

// Wraps Command{} Adapter implementation
func buildCmdFactory(cmd *Command) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return &AdapterCmd{
			cmd: cmd,
		}, nil
	}
}

func RootHelpFunc(ctx *Context, cmd *Command) cli.HelpFunc {
	// TODO: Replace with internal implementation
	// sourced from {"hashicorp/cli" | "mitchellh/cli"} HelpFunc default impl with modifications
	appName := ctx.AppName
	return func(commands map[string]cli.CommandFactory) string {
		var buf bytes.Buffer
		buf.WriteString(fmt.Sprintf("\n%s\n\n", cmd.HelpText))
		buf.WriteString("Usage:\n")
		buf.WriteString(fmt.Sprintf(
			"  %s [--version] [--help] <command> [--flags] [<args>]\n\n",
			appName))
		buf.WriteString("Available Commands:\n")

		// Get the list of keys so we can sort them, and also get the maximum
		// key length so they can be aligned properly.
		keys := make([]string, 0, len(commands))
		maxKeyLen := 0
		for key := range commands {
			if len(key) > maxKeyLen {
				maxKeyLen = len(key)
			}

			keys = append(keys, key)
		}
		sort.Strings(keys)

		for _, key := range keys {
			commandFunc, ok := commands[key]
			if !ok {
				// This should never happen since we JUST built the list of
				// keys.
				panic("command not found: " + key)
			}

			command, err := commandFunc()
			if err != nil {
				log.Printf("[ERR] cli: Command '%s' failed to load: %s",
					key, err)
				continue
			}

			key = fmt.Sprintf("%s%s", key, strings.Repeat(" ", maxKeyLen-len(key)))
			buf.WriteString(fmt.Sprintf("    %s    %s\n", key, command.Synopsis()))
		}

		return buf.String()
	}
}
