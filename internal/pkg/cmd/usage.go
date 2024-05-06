package cmd

import (
	"bytes"
	"flag"
	"fmt"
	"text/tabwriter"
	"text/template"
)

type flagUse struct {
	Name         string
	Usage        string
	DefaultValue string
}

// Returns compiled template result with any declared/expected flags
func (c *Command) flagsUsage() string {
	if c.Flags == nil {
		return ""
	}

	var strBuffer bytes.Buffer
	w := tabwriter.NewWriter(&strBuffer, 0, 8, 1, '\t', tabwriter.AlignRight)

	flagsTempl := `
Flags:{{ range . }}
  --{{ .Name }}	{{ .Usage }}	{{ if ne .DefaultValue "" }}(default "{{.DefaultValue}}"){{ end }}{{ end }}
`

	t := template.Must(template.New("flags").Parse(flagsTempl))

	valueMap := []flagUse{}
	fs := c.flagSet()
	fs.VisitAll(func(f *flag.Flag) {
		valueMap = append(valueMap, flagUse{Name: f.Name, Usage: f.Usage, DefaultValue: f.DefValue})
	})

	err := t.Execute(w, valueMap)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	w.Flush()

	return strBuffer.String()
}

// Returns compiled template result with any declared/expected arguments
func (c *Command) argumentsUsage() string {
	if len(c.Arguments.Args) <= 0 {
		return ""
	}

	var strBuffer bytes.Buffer
	w := tabwriter.NewWriter(&strBuffer, 0, 8, 1, '\t', tabwriter.AlignRight)

	argsTempl := `
Arguments:{{ range . }}
  {{ .Name }}	{{ .Description }}{{ end }}
`
	t := template.Must(template.New("flags").Parse(argsTempl))

	err := t.Execute(w, c.Arguments.Args)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	w.Flush()

	return strBuffer.String()
}
