// +build js

package dlgs

import (
//"github.com/gopherjs/gopherjs/js"
)

// File displays a file dialog, returning the selected file/directory and a bool for success.
// On Windows, if directory is true, owner should be the name of the window that owns the dialog.
func File(title, filter, owner string, directory bool) (string, bool, error) {
	return "", false, ErrNotImplemented
}

// FileMulti displays a file dialog, returning the selected files and a bool for success.
func FileMulti(title, filter string) ([]string, bool, error) {
	return []string{}, false, ErrNotImplemented
}
