package main

import (
	"strings"
	"text/template"
	"bytes"
)

func thisthis() {
        data := map[string]interface{}{
                "Name": "John Doe",
                "Age":  30,
                "Hobbies": []string{"Reading", "Hiking", "Coding"},
        }

        tmpl, err := template.New("person").Funcs(template.FuncMap{
                "indent": func(level int) string {
                        return strings.Repeat("  ", level)
                },
        }).Parse(`
{{define "person"}}
{{indent 0}}Name: {{.Name}}
{{indent 0}}Age: {{.Age}}
{{indent 0}}Hobbies:
{{range $index, $hobby := .Hobbies}}
{{indent 1}}- {{$hobby}}
{{end}}
{{end}}
`)

        if err != nil {
                fmt.Println("Error parsing template:", err)
                return
        }

        var buf bytes.Buffer
        err = tmpl.ExecuteTemplate(&buf, "person", data)
        if err != nil {
                fmt.Println("Error executing template:", err)
                return
        }

        fmt.Println(buf.String())
}
