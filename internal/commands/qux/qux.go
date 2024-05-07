package qux

import (
	"errors"
	"fmt"
	"time"

	"example.com/go-cli/internal/pkg/cmd"
)

func executeWork(resulChan chan<- struct{}) {
	fmt.Println("execution work starting...")
	time.Sleep(5 * time.Second)
	fmt.Println("execution work done...")
	resulChan <- struct{}{}
}

// Example listening to quit signals from parent ctx
func NewCmdQux(cmdCtx *cmd.Context) *cmd.Command {
	ctx := cmdCtx.ParentCtx
	ch := make(chan struct{})

	return &cmd.Command{
		Name:     "qux",
		Synopsis: "Synopsis for qux command.",
		HelpText: "Additional information about qux command",
		Run: func(c *cmd.Command, args []string) error {
			go executeWork(ch)

			select {
			// done processing
			case res := <-ch:
				fmt.Println("result chan:", res)
				return nil
			// read from chan when quit signal received
			case <-ctx.Done():
				fmt.Println("signal received")
				return errors.New("quit signal received")
			}
		},
	}
}
