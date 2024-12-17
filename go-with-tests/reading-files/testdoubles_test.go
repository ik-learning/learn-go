package main

import (
	"errors"
	"io/fs"
)

type StubFailingFS struct {
}

func (f StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}
