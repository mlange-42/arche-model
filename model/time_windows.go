//go:build windows
// +build windows

package model

import "syscall"

// initTimer adjusts time resolution for Windows.
// Does nothing on Unix systems.
// See https://github.com/golang/go/issues/44343
func initTimer() {
	winmmDLL := syscall.NewLazyDLL("winmm.dll")
	procTimeBeginPeriod := winmmDLL.NewProc("timeBeginPeriod")
	procTimeBeginPeriod.Call(uintptr(1))
}
