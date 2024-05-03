package main

import (
	"context"
	"os"

	"example.com/go-cli/internal/commands/root"
	"example.com/go-cli/internal/pkg/cmd"
	"example.com/go-cli/version"
	cli "github.com/hashicorp/cli"
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
	}

	// setup root command
	rootCmd := root.NewRootCmd(cmdCtx)
	cmdMap := cmd.NewCmdFactory(rootCmd)

	// setup hashicorp/cli cli
	cliRunner := cli.CLI{
		Args:     args,
		Commands: cmdMap,
		Version:  version.GetVersion(),
	}

	// execute command
	exitCode, err := cliRunner.Run()
	if err != nil {
		return 1
	}
	return exitCode
}
