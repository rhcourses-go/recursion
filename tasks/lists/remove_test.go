package lists

import "fmt"

func ExampleRemoveElement() {
	l1 := []int{1, 2, 3, 4, 5}

	l1 = RemoveElement(l1, 2)
	fmt.Println(l1)

	// Output:
	// [1 2 4 5]
}
