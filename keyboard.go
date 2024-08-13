package main

import (
	"fmt"
	"strings"
	"time"

	"gwintils/gwintils/util"
)

var procGetAsyncKeyState = User32.NewProc("GetAsyncKeyState")

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
