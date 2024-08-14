package mouse

import (
	"fmt"
	"unsafe"

	User32 "github.com/notaud/gwintils/base"
	"github.com/notaud/gwintils/monitor"
	"github.com/notaud/gwintils/types"
	"github.com/notaud/gwintils/util"
)

var (
	procMouseEvent   = User32.MouseEvent()
	procGetCursorPos = User32.GetCursorPos()
	procSendInput    = User32.SendInput()
)

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

func Move(x, y int32) error {
	displayWidth, displayHeight := monitor.GetDisplaySize()
	dx := util.MulDiv(x, 65536, displayWidth)
	dy := util.MulDiv(y, 65536, displayHeight)

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

func Click(button string) error {
	err := Down(button)
	if err != nil {
		return err
	}

	err = Up(button)
	if err != nil {
		return err
	}

	return nil
}

func Down(button string) error {
	var down uint32
	switch button {
	case "": // Default to Left click if nothing passed
	case "left":
		down = MOUSEEVENTF_LEFTDOWN
	case "right":
		down = MOUSEEVENTF_RIGHTDOWN
	case "middle":
		down = MOUSEEVENTF_MIDDLEDOWN
	default:
		return fmt.Errorf("unsupported button: %s", button)
	}

	err := MouseEvent(down)
	if err != nil {
		return err
	}

	return nil
}

func Up(button string) error {
	var up uint32
	switch button {
	case "": // Default to Left click if nothing passed
	case "left":
		up = MOUSEEVENTF_LEFTUP
	case "right":
		up = MOUSEEVENTF_RIGHTUP
	case "middle":
		up = MOUSEEVENTF_MIDDLEUP
	default:
		return fmt.Errorf("unsupported button: %s", button)
	}

	err := MouseEvent(up)
	if err != nil {
		return err
	}

	return nil
}

func MouseEvent(button uint32) error {
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
			DwFlags: button,
		},
	}

	ret, _, err := procSendInput.Call(
		1,
		uintptr(unsafe.Pointer(&input)),
		unsafe.Sizeof(input),
	)
	if ret == 0 {
		return fmt.Errorf("mouse_event up failed: %v", err)
	}

	return nil
}

func GetPosition() (*types.POINT, error) {
	var position types.POINT
	ret, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&position)))
	if ret == 0 {
		return nil, err
	}

	return &position, nil
}
