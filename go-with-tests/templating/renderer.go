package main

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

const (
	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)

	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

func (r *PostRenderer) Render(w io.Writer, data interface{}) error {
	if err := r.templ.ExecuteTemplate(w, "blog.gohtml", data); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.Title}}">{{.Title}}</a></li>{{end}}</ol>`
	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}
	if err := templ.Execute(w, posts); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndexV2(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.Title}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Funcs(template.FuncMap{
		"sanitiseTitle": func(title string) string {
			return strings.ToLower(strings.Replace(title, " ", "-", -1))
		},
	}).Parse(indexTemplate)

	if err != nil {
		return err
	}
	if err := templ.Execute(w, posts); err != nil {
		return err
	}
	return nil
}

func (r *PostRenderer) RenderIndexV0(w io.Writer, posts []Post) error {
	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`

	templ, err := template.New("index").Parse(indexTemplate)
	if err != nil {
		return err
	}

	if err := templ.Execute(w, posts); err != nil {
		return err
	}

	return nil
}

func (r *PostRenderer) RenderV3(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", p)
}

func (r *PostRenderer) RenderIndexV3(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

func RenderFs(w io.Writer, p Post) error {
	temp, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := temp.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
		return err
	}
	return nil
}

func RenderTemplate(w io.Writer, post Post) error {
	temp, err := template.New("blog").Parse(postTemplate)
	if err != nil {
		return err
	}
	if err := temp.Execute(w, post); err != nil {
		return err
	}
	return nil
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", post.Title, post.Description)
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "Tags: <ul>")
	if err != nil {
		return err
	}
	for _, tag := range post.Tags {
		_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}
	_, err = fmt.Fprintf(w, "</ul>")
	if err != nil {
		return err
	}
	return nil
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
