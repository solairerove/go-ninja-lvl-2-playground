package exercise2

import "fmt"

// Ex2 is a complete solution on jedi lvl 2
func Ex2() {

	x := 42 == 7
	y := 7 <= 1
	z := 7 >= 1
	a := 42 != 42
	b := 42 > 43
	c := 42 < 43

	fmt.Printf("%v\t%v\t%v\n", x, y, z)
	fmt.Printf("%v\t%v\t%v\n", a, b, c)
}
