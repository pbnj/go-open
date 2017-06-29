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
func Open(obj string) {

	if err := exec.Command(openCmd, obj).Start(); err != nil {
		fmt.Println(err)
	}

}
