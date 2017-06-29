package main

import (
	"os"

	gopen "github.com/petermbenjamin/go-open"
)

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		gopen.Open(args[0])
		return
	}
	for _, arg := range args {
		gopen.Open(arg)
	}
}
