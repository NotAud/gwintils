package monitor

import "github.com/notaud/gwintils/os"

const (
	SM_CMONITORS       = 80
	SM_XVIRTUALSCREEN  = 76
	SM_YVIRTUALSCREEN  = 77
	SM_CXVIRTUALSCREEN = 78
	SM_CYVIRTUALSCREEN = 79
)

func GetDisplays() int32 {
	return os.GetSystemMetrics(SM_CMONITORS)
}

func GetVirtualDisplaySize() (int32, int32) {
	return os.GetSystemMetrics(SM_CXVIRTUALSCREEN), os.GetSystemMetrics(SM_CYVIRTUALSCREEN)
}

func GetVirtualDisplayPosition() (int32, int32) {
	return os.GetSystemMetrics(SM_XVIRTUALSCREEN), os.GetSystemMetrics(SM_YVIRTUALSCREEN)
}
