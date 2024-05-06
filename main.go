package main

import (
	"context"
	"os"

	"example.com/go-cli/internal/commands"
	"example.com/go-cli/internal/pkg/cmd"
	"example.com/go-cli/version"
	cli "github.com/hashicorp/cli"
)

var (
	// Linker e.g. -ldflags="-X 'main.Application=app'"
	Application = "go-cli"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	args := os.Args[1:]

	// setup parent context
	// TODO: handle os signals manually for cleanup
	ctx := context.Background()

	// setup cmd package context for global fields
	cmdCtx := &cmd.Context{
		ParentCtx: ctx,
		AppName:   Application,
	}

	// setup root command
	rootCmd := commands.NewCmdRoot(cmdCtx)
	cmdMap := cmd.NewCmdFactory(rootCmd)

	// setup hashicorp/cli cli
	cliRunner := cli.CLI{
		Name:     Application,
		Args:     args,
		Commands: cmdMap,
		Version:  version.GetVersion(),
		HelpFunc: cmd.RootHelpFunc(cmdCtx, rootCmd),
	}

	// execute command
	exitCode, err := cliRunner.Run()
	if err != nil {
		return 1
	}
	return exitCode
}
