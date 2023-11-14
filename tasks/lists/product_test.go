package lists

import "fmt"

func ExampleProduct() {
	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{2, 4, 6, 8, 10}
	l3 := []int{}
	l4 := []int{2, 3, 2}

	fmt.Println(Product(l1))
	fmt.Println(Product(l2))
	fmt.Println(Product(l3))
	fmt.Println(Product(l4))

	// Output:
	// 120
	// 3840
	// 1
	// 12
}
