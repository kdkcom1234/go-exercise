package pointer

import (
	"fmt"
)

func SetPointerVar() {
	var a int
	p := &a
	*p = 50

	fmt.Println(*p, a)
}

func SetTo100() {
	var num int
	p := &num
	*p = 100

	fmt.Println(*p, num)
}
