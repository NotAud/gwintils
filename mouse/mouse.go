package mouse

import (
	"fmt"
	"unsafe"

	User32 "github.com/notaud/gwintils/base"
	"github.com/notaud/gwintils/monitor"
)

var (
	procMouseEvent   = User32.MouseEvent()
	procGetCursorPos = User32.GetCursorPos()
	procSendInput    = User32.SendInput()
)

type Input struct {
	Type uint32
	Mi   MouseInput
}

type MouseInput struct {
	Dx          int32
	Dy          int32
	MouseData   uint32
	DwFlags     uint32
	Time        uint32
	DwExtraInfo uintptr
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
	vScreenWidth, vScreenHeight := monitor.GetVirtualDisplaySize()
	vScreenLeft, vScreenTop := monitor.GetVirtualDisplayPosition()

	relativeX := x - vScreenLeft
	relativeY := y - vScreenTop

	dx := (float64(relativeX) * 65535) / float64(vScreenWidth)
	dy := (float64(relativeY) * 65535) / float64(vScreenHeight)

	fmt.Println(relativeX, relativeY)

	const dwFlags = MOUSEEVENTF_MOVE | MOUSEEVENTF_VIRTUALDESK | MOUSEEVENTF_ABSOLUTE
	input := Input{
		Type: INPUT_MOUSE,
		Mi: MouseInput{
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
	input := Input{
		Type: INPUT_MOUSE,
		Mi: MouseInput{
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

type Test struct {
	X, Y int32
}

func GetPosition() (*Test, error) {
	var position Test
	ret, _, err := procGetCursorPos.Call(uintptr(unsafe.Pointer(&position)))
	if ret == 0 {
		return nil, err
	}

	return &position, nil
}
