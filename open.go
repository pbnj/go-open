package open

import (
	"fmt"
	"os/exec"
	"runtime"
)

var openCmd string

func init() {
	if runtime.GOOS == "windows" {
		openCmd = "start"
	} else if runtime.GOOS == "darwin" {
		openCmd = "open"
	} else if runtime.GOOS == "linux" {
		openCmd = "xdg-open"
	}
}

// Open will open things in their default applications
func Open(things []string) {

	for _, thing := range things {

		if err := exec.Command(openCmd, thing).Start(); err != nil {
			fmt.Println(err)
		}

	}

}
