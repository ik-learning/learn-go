package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRunner(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := new(bytes.Buffer)
		err := Render(buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>hello world</h1>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := new(bytes.Buffer)
		err := Render(buf, aPost)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<h1>hello world</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestRenderV2(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}

		if err := Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}

		fmt.Println(buf.String())

		approvals.VerifyString(t, buf.String())
	})
}

func TestRenderV3(t *testing.T) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)
	postRenderer, err := NewPostRenderer()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := new(bytes.Buffer)
		if err := postRenderer.Render(buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndexV2(&buf, posts); err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndexV3(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

// BenchmarkRender-16    	 2846858	       417.2 ns/op
func BenchmarkRender(b *testing.B) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Render(io.Discard, aPost)
	}
}

func BenchmarkRenderV2(b *testing.B) {
	var (
		aPost = Post{
			Title:       "hello world",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
