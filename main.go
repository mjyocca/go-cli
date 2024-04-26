package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, []any{"Hello World"}...)
}
