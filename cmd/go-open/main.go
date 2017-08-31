package main

import (
	"fmt"
	"os"

	gopen "github.com/petermbenjamin/go-open"
)

func main() {
	for _, arg := range os.Args[1:] {
		if err := gopen.Open(arg); err != nil {
			fmt.Println(err)
		}
	}
}
