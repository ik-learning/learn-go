package main

import (
	"reflect"
	// "io/fs"
	"testing/fstest"

	"testing"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}

	posts, _ := NewPostsFromFS(fs)
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func TestNewBlogPostsV2(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

// func TestNewBlogPostsV3(t *testing.T) {
// 	const (
// 		firstBody  = `Title: Post 1 Description: Description 1 Tags: tdd, go`
// 		secondBody = `Title: Post 2 Description: Description 2 Tags: rust, borrow-checker`
// 	)
// 	fs := fstest.MapFS{
// 		"hello world.md":  {Data: []byte(firstBody)},
// 		"hello-world2.md": {Data: []byte(secondBody)},
// 	}
//
// 	posts, err := NewPostsFromFs(fs)
// 	assertNoError(t, err)
//
// 	assertPost(t, posts[0], &Post{
// 		Title:       "Post 1",
// 		Description: "Description 1",
// 		Tags:        []string{"tdd", "go"},
// 	})
// }

func TestNewBlogPostsV5(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := NewPostsFromFS(fs)

	assertNoError(t, err)

	assertPostsLength(t, posts, fs)

	assertPost(t, posts[0], Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
