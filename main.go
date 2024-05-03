package main

import (
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

	// setup root command
	rootCmd := root.NewRootCmd()
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
