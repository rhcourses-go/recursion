package calc

import "fmt"

func ExampleMult() {
	fmt.Println(Mult(3, 4))
	fmt.Println(Mult(21, 2))
	fmt.Println(Mult(2, 0))
	fmt.Println(Mult(2, 1))

	// Output:
	// 12
	// 42
	// 0
	// 2
}
