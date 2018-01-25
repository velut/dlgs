// +build darwin,!linux,!windows,!js

package dlgs

import (
	"os/exec"
	"strings"
	"syscall"
)

// Entry displays input dialog, returning the entered value and a bool for success.
func Entry(title, text, defaultText string) (string, bool, error) {
	osa, err := exec.LookPath("osascript")
	if err != nil {
		return "", false, err
	}

	o, err := exec.Command(osa, "-e", `set T to text returned of (display dialog "`+text+`" with title "`+title+`" default answer "`+defaultText+`")`).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			return "", ws.ExitStatus() == 0, nil
		}
	}

	ret := true
	out := strings.TrimSpace(string(o))
	if out == "" {
		ret = false
	}

	return out, ret, err
}
