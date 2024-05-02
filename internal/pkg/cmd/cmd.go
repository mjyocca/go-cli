package cmd

import "github.com/hashicorp/cli"

func BuildCmdMap(cmd *Command) map[string]cli.CommandFactory {
	m := map[string]cli.CommandFactory{}

	// build alias slice to iterate over
	aliases := []string{cmd.Name}
	for _, alias := range cmd.Aliases {
		aliases = append(aliases, alias)
	}

	// loop over name +aliases
	for _, name := range aliases {
		m[name] = func() (cli.Command, error) {
			return &commandCli{
				cmd: cmd,
			}, nil
		}
	}

	return m
}
