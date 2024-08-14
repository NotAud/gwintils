package os

import User32 "github.com/notaud/gwintils/base"

var procGetSystemMetrics = User32.GetSystemMetrics()

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(nIndex))
	return int(ret)
}
