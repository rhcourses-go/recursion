package lists

import "fmt"

func ExampleEmpty() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{}

	fmt.Println(Empty(l1))
	fmt.Println(Empty(l2))

	// Output:
	// false
	// true
}
