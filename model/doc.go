// Package model provides a wrapper around the Arche ECS world
// that helps with rapid prototyping and model development.
package model

// init initializes the package.
// It adjusts time resolution for Windows.
// See this issue: https://github.com/golang/go/issues/44343
func init() {
	initTimer()
}
