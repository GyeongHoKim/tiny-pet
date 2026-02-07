//go:build !debug

package main

// debugPrint is a no-op in release builds to reduce firmware size (no println/formatting).
func debugPrint(...interface{}) {}
