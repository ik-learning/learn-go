package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/jedib0t/go-pretty/v6/table"
)

type Table struct {
	Table table.Writer
}

type Flag struct {
	Name        string
	Description string
	EnvVar      string
}
type Flags struct {
	Flag          []Flag
	MaxNameLength int
	MaxDescLength int
}

// {{repeat "*" 10}}
const markdownTemplate = `
# Flags

| Command{{ computeIndent "command" (len "Command") | indent }} | Description{{ computeIndent "desc" (len "Description") | indent }} |
| :-{{ repeat "command" "-"}} | :-{{ repeat "desc" "-"}} |
{{- range . }}
| {{ .Name }}{{ computeIndent "command" (len .Name) | indent }} | {{ .Description }}{{ computeIndent "desc" (len .Description) | indent }} | {{- end -}}
`

func (f *Flags) GenerateMarkdownTable() (string, error) {
	tmpl, err := template.New("markdown").Funcs(template.FuncMap{
		"indent": func(level int) string {
			return strings.Repeat(" ", level)
		},
		"computeIndent": func(column string, length int) int {
			var maximum int
			if column == "command" {
				maximum = f.MaxNameLength
			} else {
				maximum = f.MaxDescLength
			}
			return maximum - length
		},
		"replace": strings.ReplaceAll,
		"join":    strings.Join,
		"repeat": func(column string, s string) string {
			var maximum int
			if column == "command" {
				maximum = f.MaxNameLength
			} else {
				maximum = f.MaxDescLength
			}
			return strings.Repeat(s, maximum-2)
		},
	}).Parse(markdownTemplate)
	if err != nil {
		return "", err
	}

	var table bytes.Buffer
	err = tmpl.Execute(&table, f.Flag)
	if err != nil {
		return "", err
	}

	return table.String(), nil
}
func (f *Flags) AddFlag(flag Flag) {
	if len(flag.Name) > f.MaxNameLength {
		f.MaxNameLength = len(flag.Name)
	}
	if len(flag.Description) > f.MaxDescLength {
		f.MaxDescLength = len(flag.Description)
	}
	f.Flag = append(f.Flag, flag)
}

func writeToFile(filename string, content string) error {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		fmt.Errorf("failed to create file: %v", fileErr)
	}
	defer file.Close()

	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		fmt.Errorf("Failed to write to file: %s", filename)
	}

	return nil
}

func main() {
	cfg := Config{}
	app, err := cfg.ParseFlags([]string{"--help-markdown"})
	if err != nil {
		fmt.Println("error", err)
	}
	kingpinFlags := app.Model().Flags

	var f []Flag
	fff := Flags{
		Flag: f,
	}

	for _, flag := range kingpinFlags {
		// do not include helpers and completion flags
		if strings.Contains(flag.Name, "help") || strings.Contains(flag.Name, "completion-") {
			continue
		}
		flagString := ""
		flagName := flag.Name
		if flag.IsBoolFlag() {
			flagName = "[no-]" + flagName
		}
		flagString += fmt.Sprintf("--%s", flagName)

		if !flag.IsBoolFlag() {
			flagString += fmt.Sprintf("=%s", flag.FormatPlaceHolder())
		}
		fff.AddFlag(Flag{Name: fmt.Sprintf("`%s`", flagString), Description: flag.HelpWithEnvar()})
	}

	table, err := fff.GenerateMarkdownTable()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	fmt.Println(table)

	_ = writeToFile("cli/kingpin/flags.md", table)
}
