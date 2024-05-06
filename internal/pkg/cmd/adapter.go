package cmd

import (
	"fmt"

	"github.com/hashicorp/cli"
)

// verify AdapterCmd satisfies hashicorp/cli interface(s)
var _ cli.Command = (*AdapterCmd)(nil)
var _ cli.CommandHelpTemplate = (*AdapterCmd)(nil)

// Adapter serving as a mediator between cmd.Command struct{} and hashicorp cli.Command interface{}
type AdapterCmd struct {
	cmd *Command
}

// implements cli.Command interface{}
func (c *AdapterCmd) Run(args []string) int {
	return c.cmd.run(args)
}

func (c *AdapterCmd) Help() string {
	return c.cmd.HelpText
}

func (c *AdapterCmd) Synopsis() string {
	return c.cmd.Synopsis
}

// implements cli.CommandTemplate interface{}
func (c *AdapterCmd) HelpTemplate() string {
	flags := c.cmd.flagsUsage()
	args := c.cmd.argumentsUsage()

	// TODO: build subcommand template and replace this
	prepare := `
{{.Help}}{{if gt (len .Subcommands) 0}}

Available Commands:
{{- range $value := .Subcommands }}
    {{ $value.NameAligned }}    {{ $value.Synopsis }}{{ end }}
{{- end }}
%s%s
Use "<app-name> [command] --help" for more information about a command.
`
	// TODO: allow each cmd to access appname
	prepare = fmt.Sprintf(prepare, flags, args)

	return prepare
}
