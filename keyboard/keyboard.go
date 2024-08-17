package keyboard

import (
	"fmt"
	"strings"

	User32 "github.com/notaud/gwintils/base"
	"github.com/notaud/gwintils/util"
)

const KEY_PRESSED_BITMASK uintptr = 0x8000

var procGetAsyncKeyState = User32.GetAsyncKeyState()

func GetAsyncKeyState(vKey uintptr) bool {
	state, _, _ := procGetAsyncKeyState.Call(vKey)
	return state&KEY_PRESSED_BITMASK != 0
}

func VirtualKey(key string) (uintptr, error) {
	vk, ok := util.KeyMap[strings.ToLower(key)]
	if !ok {
		return 0, fmt.Errorf("unsupported key: %s", key)
	}
	return vk, nil
}
