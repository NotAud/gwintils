package monitor

import "github.com/notaud/gwintils/os"

const (
	SM_CMONITORS       = 80
	SM_CXVIRTUALSCREEN = 78
	SM_CYVIRTUALSCREEN = 79
)

func GetDisplays() int32 {
	return os.GetSystemMetrics(SM_CMONITORS)
}

func GetDisplaySize() (int32, int32) {
	return os.GetSystemMetrics(SM_CXVIRTUALSCREEN), os.GetSystemMetrics(SM_CYVIRTUALSCREEN)
}
