//go:build !windows
// +build !windows

package model

// initTimer adjusts time resolution for Windows.
// Does nothing on Unix systems.
// See this issue: https://github.com/golang/go/issues/44343
func initTimer() {
}
