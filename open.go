package open

import (
	"os/exec"
	"runtime"

	"github.com/pkg/errors"
)

// Open will open resources (e.g. files, URLs ...etc) in their default applications
func Open(obj string) error {
	switch runtime.GOOS {
	case "darwin":
		if err := exec.Command("open", obj).Start(); err != nil {
			return errors.Errorf("could not open resource: %s", err)
		}
	case "linux":
		if err := exec.Command("xdg-open", obj).Start(); err != nil {
			return errors.Errorf("could not open resource: %s", err)
		}
	case "windows":
		if err := exec.Command("cmd", "/c", "start", obj).Start(); err != nil {
			return errors.Errorf("could not open resource: %s", err)
		}
	default:
		return errors.Errorf("could not identify OS")
	}
	return nil
}
