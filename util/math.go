package util

func MulDiv(number, numerator, denominator int32) int32 {
	return (number*numerator + (denominator >> 1)) / denominator
}
