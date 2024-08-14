package keyboard

import (
	"fmt"
	"strings"
	"time"

	User32 "github.com/notaud/gwintils/base"
	"github.com/notaud/gwintils/util"
)

var procGetAsyncKeyState = User32.GetAsyncKeyState()

func ListenKey(key string) error {
	vk, ok := util.KeyMap[strings.ToLower(key)]
	if !ok {
		return fmt.Errorf("unsupported key: %s", key)
	}

	for {
		state, _, _ := procGetAsyncKeyState.Call(vk)
		if state&0x8000 != 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	return nil
}
