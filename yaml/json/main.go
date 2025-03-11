package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Changes struct {
	// Records that need to be created
	Create []*Endpoint `json:"create"`
	// Records that need to be updated (current data)
	UpdateOld []*Endpoint
	// Records that need to be updated (desired data)
	UpdateNew []*Endpoint
	// Records that need to be deleted
	Delete []*Endpoint
}

type Endpoint struct {
	Name string
}

var (
	input = `
{
  "create": [
    {
      "Name": "example.com"
    },
    {
      "Name": "example.org"
    }
  ],
  "updateold": [
    {
      "Name": "old.example.com"
    },
    {
      "naMe": "paradox.example.com"
    }
  ],
  "updatenew": [
    {
      "Name": "new.example.com"
    }
  ],
  "Delete": [
    {
      "Name": "delete.example.com"
    }
  ]
}
`
)

func main() {
	fmt.Println("Hello, playground")

	changes := Changes{
		Create: []*Endpoint{
			{Name: "example.com"},
			{Name: "example.org"},
		},
		UpdateOld: []*Endpoint{
			{Name: "old.example.com"},
		},
		UpdateNew: []*Endpoint{
			{Name: "new.example.com"},
		},
		Delete: []*Endpoint{
			{Name: "delete.example.com"},
		},
	}

	// Print the changes object to verify
	data, err := json.MarshalIndent(changes, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling changes:", err)
		return
	}

	fmt.Println(string(data))

	var ch Changes

	reader := io.NopCloser(strings.NewReader(input))
	if err := json.NewDecoder(reader).Decode(&ch); err != nil {
		return
	}
	fmt.Println("DECODED")
	fmt.Println("Create:")
	for _, e := range ch.Create {
		fmt.Println(e.Name)
	}
	fmt.Println("UpdateOld:")
	for _, e := range ch.UpdateOld {
		fmt.Println(e.Name)
	}
	fmt.Println("Delete:")
	for _, e := range ch.Delete {
		fmt.Println(e.Name)
	}
}
