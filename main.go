package main

import (
	"golang.org/x/sys/windows"
)

var User32 = windows.NewLazySystemDLL("user32.dll")
