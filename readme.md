# Go-CLI

## Example usage

```go
//...
cmd := &cmd.Command{
	Name:     "<command>",
	Synopsis: "<synopsis>",
	HelpText: "<help>",
	Flags: func(f *flag.FlagSet) {
		f.BoolVar(&opts.enabled, "enabled", false, "...")
	},
	Arguments: cmd.Arguments{
		Args: []cmd.Argument{
			{
				Name:        "<arg-name>",
				Description: "<arg-description>",
			},
		},
		Validate: cmd.ExactArgs(1),
	},
	Run: func(c *cmd.Command, args []string) error {
		//...
	},
}
```
