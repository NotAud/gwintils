package main

import (
	"fmt"
	"unsafe"
)

var procSendInput = User32.NewProc("SendInput")

type Inputer interface {
	isInput()
}

// Try to fix later

func SendInput(input Inputer) error {
	fmt.Println(input, unsafe.Sizeof(input))

	_, _, err := procSendInput.Call(
		1,
		uintptr(unsafe.Pointer(&input)),
		unsafe.Sizeof(input),
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
