package cmd

import "fmt"

// function used to validate a command receives expected/valid arguments.
type ValidateArgsFunc func(c *Command, args []string) error

type Arguments struct {
	Args     []Argument
	Validate ValidateArgsFunc
}

type Argument struct {
	// Name of the argument
	Name string
	// Description of the argument
	Description string
}

func (a *Arguments) validateFunc() ValidateArgsFunc {
	if a.Validate != nil {
		return a.Validate
	}

	// TODO if no validate specified, check Args
	return func(cmd *Command, args []string) error {
		return nil
	}
}

func NoArgs(cmd *Command, args []string) error {
	if len(args) > 0 {
		return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
	}
	return nil
}

func MinimumNArgs(n int) ValidateArgsFunc {
	return func(cmd *Command, args []string) error {
		if len(args) < n {
			return fmt.Errorf("requires at least %d arg(s), only received %d", n, len(args))
		}
		return nil
	}
}

func MaximumNArgs(n int) ValidateArgsFunc {
	return func(cmd *Command, args []string) error {
		if len(args) > n {
			return fmt.Errorf("accepts at most %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}

func RangeArgs(min int, max int) ValidateArgsFunc {
	return func(cmd *Command, args []string) error {
		if len(args) < min || len(args) > max {
			return fmt.Errorf("accepts between %d and %d arg(s), received %d", min, max, len(args))
		}
		return nil
	}
}

func ExactArgs(n int) ValidateArgsFunc {
	return func(cmd *Command, args []string) error {
		if len(args) != n {
			return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
		}
		return nil
	}
}
