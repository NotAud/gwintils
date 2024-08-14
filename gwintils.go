package main

import (
	"golang.org/x/sys/windows"
)

const (
	Version = "v0.0.1"
)

var User32 = windows.NewLazySystemDLL("user32.dll")
