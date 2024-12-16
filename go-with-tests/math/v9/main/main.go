package main

import (
	"os"
	"time"

	clockface "github.com/ik-learning/learn-go/go-with-tests/math/v9"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
