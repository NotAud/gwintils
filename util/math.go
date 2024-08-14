package util

import (
	User32 "github.com/notaud/gwintils/base"
)

var procMulDiv = User32.MulDiv()

func MulDiv(number, numerator, denominator int32) int32 {
	ret, _, _ := procMulDiv.Call(
		uintptr(number),
		uintptr(numerator),
		uintptr(denominator),
	)
	return int32(ret)
}
