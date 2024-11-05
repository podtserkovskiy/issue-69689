package testdata

import "C"

func Foo(number C.int) int {
	return int(number)
}
