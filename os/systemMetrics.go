package os

import (
	"unsafe"

	User32 "github.com/notaud/gwintils/base"
	"github.com/notaud/gwintils/types"
)

var (
	procGetSystemMetrics              = User32.GetSystemMetrics()
	procSetProcessDpiAwarenessContext = User32.SetProcessDpiAwarenessContext()
)

var (
	DPI_AWARENESS_CONTEXT_UNAWARE              = NewHWND() - 1
	DPI_AWARENESS_CONTEXT_SYSTEM_AWARE         = NewHWND() - 2
	DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE    = NewHWND() - 3
	DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2 = NewHWND() - 4
	DPI_AWARENESS_CONTEXT_UNAWARE_GDISCALED    = NewHWND() - 5
)

func GetSystemMetrics(nIndex int32) int32 {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(nIndex))
	return int32(ret)
}

func SetProcessDpiAwarenessContext(dpiAwarenessLevel uintptr) error {
	ret, _, err := procSetProcessDpiAwarenessContext.Call(dpiAwarenessLevel)
	if ret != 0 {
		return err
	}

	return nil
}

func NewHWND() types.HWND {
	return types.HWND(uintptr(unsafe.Pointer(nil)))
}
