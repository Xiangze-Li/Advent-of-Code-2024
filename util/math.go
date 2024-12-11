package util

import "slices"

type Integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~uintptr
}

var digitBase = [...]uint64{
	1, 10, 100,
	1_000, 10_000, 100_000,
	1_000_000, 10_000_000, 100_000_000,
	1_000_000_000, 10_000_000_000, 100_000_000_000,
	1_000_000_000_000, 10_000_000_000_000, 100_000_000_000_000,
	1_000_000_000_000_000, 10_000_000_000_000_000, 100_000_000_000_000_000,
	1_000_000_000_000_000_000, 10_000_000_000_000_000_000,
}

// CountDigits returns the number of decimal digits in n.
func CountDigits[Int Integer](n Int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
	}
	i, f := slices.BinarySearch(digitBase[:], uint64(n))
	if f {
		return i + 1
	}
	return i
}
