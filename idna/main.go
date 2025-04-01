package main

import (
	"fmt"

	"golang.org/x/net/idna"
)

const (
	// IDNA2008 is the IDNA2008 profile.
	domain1 = "xn--eckh0ome.xn--q9jyb4c"
	domain2 = "エイミー.みんな"
)

func main() {

	var hostname = "*.foo.com"
	name, err := idna.Lookup.ToUnicode(hostname)
	if err != nil {
		fmt.Println("test")
		fmt.Errorf("failed to convert hostname '%s' to its Unicode form: %v", hostname, err)
		name = hostname
	}
	fmt.Println(name)
}
