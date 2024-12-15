package main

import (
	"fmt"
	"regexp"
)

func main() {
	var re = regexp.MustCompile("`runAsNonRoot` is set to `false` but `runAsUser` is not set to `0` (root)")
	var str = `(root)`

	if len(re.FindStringIndex(str)) > 0 {
		fmt.Println(re.FindString(str), "found at index", re.FindStringIndex(str)[0])
	}
}
