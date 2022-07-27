package go_bugs

import (
	"fmt"
)

func ExampleDeferReturn() {
	fmt.Println(deferReturn())

	// output:
	// 1
}

func ExampleDeferReturn2() {
	fmt.Println(deferReturn2())

	// output:
	// 2
}
