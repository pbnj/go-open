package main

import (
	"os"

	open "github.com/petermbenjamin/go-open"
)

func main() {
	open.Open(os.Args[1:])
}
