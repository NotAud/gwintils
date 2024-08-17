package types

type Point struct {
	X, Y int32
}

type Rect struct {
	Left, Top, Right, Bottom int32
}

type HWND = uintptr
