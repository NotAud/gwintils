package main

import (
	"fmt"
	"time"
	"unsafe"
)

var procMouseEvent = User32.NewProc("mouse_event")

type MouseInput struct {
	Type uint32
	Mi   struct {
		Dx          int32
		Dy          int32
		MouseData   uint32
		DwFlags     uint32
		Time        uint32
		DwExtraInfo uintptr
	}
}

func (MouseInput) isInput() {}

const (
	INPUT_MOUSE             = 0
	MOUSEEVENTF_MOVE        = 0x0001
	MOUSEEVENTF_ABSOLUTE    = 0x8000
	MOUSEEVENTF_LEFTDOWN    = 0x0002
	MOUSEEVENTF_LEFTUP      = 0x0004
	MOUSEEVENTF_RIGHTDOWN   = 0x0008
	MOUSEEVENTF_RIGHTUP     = 0x0010
	MOUSEEVENTF_MIDDLEDOWN  = 0x0020
	MOUSEEVENTF_MIDDLEUP    = 0x0040
	MOUSEEVENTF_VIRTUALDESK = 0x4000
)

func MouseMove(x, y int32) error {
	displayWidth, displayHeight := GetDisplaySize()
	dx := float64(x) * 65535 / float64(displayWidth)
	dy := float64(y) * 65535 / float64(displayHeight)

	const dwFlags = MOUSEEVENTF_MOVE | MOUSEEVENTF_VIRTUALDESK | MOUSEEVENTF_ABSOLUTE
	input := MouseInput{
		Type: INPUT_MOUSE,
		Mi: struct {
			Dx          int32
			Dy          int32
			MouseData   uint32
			DwFlags     uint32
			Time        uint32
			DwExtraInfo uintptr
		}{
			Dx:      int32(dx),
			Dy:      int32(dy),
			DwFlags: dwFlags,
		},
	}

	ret, _, err := procSendInput.Call(
		1,
		uintptr(unsafe.Pointer(&input)),
		unsafe.Sizeof(input),
	)

	if ret == 0 {
		return err
	}

	return nil
}

func MouseButton(button string) error {
	var down, up uintptr
	switch button {
	case "left":
		down, up = MOUSEEVENTF_LEFTDOWN, MOUSEEVENTF_LEFTUP
	case "right":
		down, up = MOUSEEVENTF_RIGHTDOWN, MOUSEEVENTF_RIGHTUP
	case "middle":
		down, up = MOUSEEVENTF_MIDDLEDOWN, MOUSEEVENTF_MIDDLEUP
	default:
		return fmt.Errorf("unsupported button: %s", button)
	}

	_, _, err := procMouseEvent.Call(down, 0, 0, 0, 0)
	if err != nil && err.Error() != "The operation completed successfully." {
		return fmt.Errorf("mouse_event down failed: %v", err)
	}

	time.Sleep(1 * time.Millisecond)

	_, _, err = procMouseEvent.Call(up, 0, 0, 0, 0)
	if err != nil && err.Error() != "The operation completed successfully." {
		return fmt.Errorf("mouse_event up failed: %v", err)
	}

	return nil
}
