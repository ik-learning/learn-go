package main

import "fmt"

const (
	RecordTypeAAAA = "AAAA"
	RecordTypeA    = "A"
)

type Targets []string

type Endpoint struct {
	RecordType string
	IP         string
	Targets
}

func main() {
	opt1()
	opt2()
}

func opt1() {
	fmt.Println("opt1")
	endpoint1 := &Endpoint{
		RecordType: RecordTypeAAAA,
		IP:         "2001:4860:4860::8888",
		Targets:    []string{"a", "b", "c"},
	}
	endpoint2 := *endpoint1
	endpoint2.RecordType = RecordTypeA
	endpoint2.IP = "100.100.1.1"
	endpoint2.Targets = []string{"d", "e", "f"}

	fmt.Println("first:", endpoint1)
	fmt.Println("second", endpoint2)
	fmt.Println("-----")
}

func opt2() {
	fmt.Println("opt2")
	endpoint1 := &Endpoint{
		RecordType: RecordTypeAAAA,
		IP:         "2001:4860:4860::8888",
	}
	endpoint2 := endpoint1
	endpoint2.RecordType = RecordTypeA
	endpoint2.IP = "100.100.1.1"

	fmt.Println("first:", endpoint1)
	fmt.Println("second", endpoint2)
	fmt.Println("-----")
}
