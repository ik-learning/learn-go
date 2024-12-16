// Writes an SVG clockface of the current time to Stdout.
package main

import (
	"os"
	"time"

	"github.com/ik-learning/learn-go/go-with-tests/math/v10/svg"
)

func main() {
	t := time.Now()
	svg.Write(os.Stdout, t)
}
