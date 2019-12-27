package main

import (
	"fmt"
	"os"

	open "github.com/pbnj/go-open"
)

func main() {
	for _, arg := range os.Args[1:] {
		if err := open.Open(arg); err != nil {
			fmt.Println(err)
		}
	}
}
