//go:build debug

package main

func debugPrint(a ...interface{}) {
	println(a...)
}
