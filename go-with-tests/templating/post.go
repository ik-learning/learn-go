package main

import (
	"html/template"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Body        string
	Tags        []string
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}
