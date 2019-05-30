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
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("runtime.GOOS %s is not supported", runtime.GOOS)
	}
}
