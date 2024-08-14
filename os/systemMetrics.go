package os

import User32 "github.com/notaud/gwintils/base"

var procGetSystemMetrics = User32.GetSystemMetrics()

func GetSystemMetrics(nIndex int32) int32 {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(nIndex))
	return int32(ret)
}
