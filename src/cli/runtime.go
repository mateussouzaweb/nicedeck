package cli

import (
	"runtime"
)

// Check if running system is ARM 64 bits
func IsArm64() bool {
	return runtime.GOARCH == "arm"
}

// Check if running system is AMD 64 bits
func IsAmd64() bool {
	return runtime.GOARCH == "amd64"
}

// Check if running system is Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// Check if running system is MacOS
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// Check if running system is Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// Return the best string based on running system
func ArchVariant(amd64 string, arm64 string) string {
	if IsAmd64() {
		return amd64
	}

	return arm64
}
