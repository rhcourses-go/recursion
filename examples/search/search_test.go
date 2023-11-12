package search

import "fmt"

// ExampleFind zeigt die Verwendung von Find.
func ExampleFind() {
	list := []int{3, 2, 5, 3, 5, 4}

	fmt.Println(Find(list, 3))
	fmt.Println(Find(list, 5))
	fmt.Println(Find(list, 10))

	// Output:
	// 0
	// 2
	// -1
}

// ExampleFindTail zeigt die Verwendung von FindTail.
func ExampleFindTail() {
	list := []int{3, 2, 5, 3, 5, 4}

	fmt.Println(FindTail(list, 3))
	fmt.Println(FindTail(list, 5))
	fmt.Println(FindTail(list, 10))

	// Output:
	// 0
	// 2
	// -1
}

// ExampleFindSorted zeigt die Verwendung von FindSorted.
func ExampleFindSorted() {
	list := []int{2, 3, 4, 10, 20, 30}

	fmt.Println(FindSorted(list, 3))
	fmt.Println(FindSorted(list, 10))
	fmt.Println(FindSorted(list, 40))

	// Output:
	// 1
	// 3
	// -1
}
