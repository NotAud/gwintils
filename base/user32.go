package User32

import "golang.org/x/sys/windows"

var User32 = windows.NewLazySystemDLL("user32.dll")

func SetProcessDpiAwarenessContext() *windows.LazyProc {
	return User32.NewProc("SetProcessDpiAwarenessContext")
}

func GetAsyncKeyState() *windows.LazyProc {
	return User32.NewProc("GetAsyncKeyState")
}

func MouseEvent() *windows.LazyProc {
	return User32.NewProc("mouse_event")
}

func GetCursorPos() *windows.LazyProc {
	return User32.NewProc("GetCursorPos")
}

func GetSystemMetrics() *windows.LazyProc {
	return User32.NewProc("GetSystemMetrics")
}

func SendInput() *windows.LazyProc {
	return User32.NewProc("SendInput")
}
