package main

const (
	SM_CMONITORS       = 80
	SM_CXVIRTUALSCREEN = 78
	SM_CYVIRTUALSCREEN = 79
)

func GetDisplays() int {
	return GetSystemMetrics(SM_CMONITORS)
}

func GetDisplaySize() (int, int) {
	return GetSystemMetrics(SM_CXVIRTUALSCREEN), GetSystemMetrics(SM_CYVIRTUALSCREEN)
}
