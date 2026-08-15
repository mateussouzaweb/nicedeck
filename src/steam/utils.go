package steam

import (
	"strings"
)

// Ensure executable by adding special wrappers when necessary
func EnsureExec(runtime string, exec string) string {
	if runtime == "flatpak" {
		return "/usr/bin/flatpak-spawn --host " + CleanExec(exec)
	}
	return CleanExec(exec)
}

// Clean executable by removing special wrappers
func CleanExec(exec string) string {
	exec = strings.Replace(exec, "/usr/bin/flatpak-spawn --host", "", 1)
	exec = strings.Trim(exec, " ")
	return exec
}
