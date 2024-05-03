package cmd

type Command struct {
	Name  string
	Short string
	Long  string

	Aliases []string

	Execute func(c *Command, args []string) error
}
