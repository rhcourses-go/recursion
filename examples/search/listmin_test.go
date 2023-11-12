package search

import "fmt"

// ExampleFindMin zeigt die Verwendung von FindMin.
func ExampleFindMin() {
	list := []int{3, 2, 5, 3, 5, 4}

	fmt.Println(FindMin(list))

	// Output:
	// 2
}
