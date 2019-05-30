package openurl

import (
	"fmt"
	"os/exec"
	"runtime"
)

// OpenWithBrowser opens url with default browser.
// It supports to OS are linux, windows and darwin(MacOS).
// If runtime.GOOS is not supported OS, this returns error.
func OpenWithBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		exec.Command("xdg-open", url).Start()
	case "windows":
		exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("runtime.GOOS %s is not supported", runtime.GOOS)
	}
	return err
}
