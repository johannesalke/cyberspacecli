//go:build windows

package main

import (
	"golang.org/x/sys/windows"
	"os"
)

func EnableANSI() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	windows.GetConsoleMode(stdout, &mode)
	windows.SetConsoleMode(stdout, mode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)

	//this
}
