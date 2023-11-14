package lists

import "fmt"

func ExampleListsEqual() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{1, 2, 3, 5, 6}
	l3 := []int{1, 2, 3}

	fmt.Println(ListsEqual(l1, l1))
	fmt.Println(ListsEqual(l1, l2))
	fmt.Println(ListsEqual(l1, l3))

	// Output:
	// true
	// false
	// false
}
