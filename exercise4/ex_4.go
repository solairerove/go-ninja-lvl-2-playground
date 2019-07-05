package exercise4

import "fmt"

const (
	x int = 42
)

// Ex4 is a complete solution on jedi lvl 2
func Ex4() {

	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	y := x << 1

	fmt.Printf("%d\t%b\t%#x\n", y, y, y)
}
