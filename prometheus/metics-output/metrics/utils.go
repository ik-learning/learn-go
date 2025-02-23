package metrics

import (
	"fmt"
	"os"
	"strings"
	"text/template"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func writeToFile(filename string, content string) error {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		_ = fmt.Errorf("failed to create file: %v", fileErr)
	}
	defer file.Close()
	_, writeErr := file.WriteString(content)
	if writeErr != nil {
		_ = fmt.Errorf("failed to write to file: %s", filename)
	}
	return nil
}

// funcMap returns a mapping of all of the functions that Engine has.
func funcMap() template.FuncMap {
	return template.FuncMap{
		"backtick": func(times int) string {
			return strings.Repeat("`", times)
		},
		"capitalize": cases.Title(language.English, cases.Compact).String,
	}
}
