package main

var procGetSystemMetrics = User32.NewProc("GetSystemMetrics")

func GetSystemMetrics(nIndex int) int {
	ret, _, _ := procGetSystemMetrics.Call(uintptr(nIndex))
	return int(ret)
}
