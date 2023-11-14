package strings

import "fmt"

func ExampleStartsWith() {
	fmt.Println(StartsWith("abc", "a"))
	fmt.Println(StartsWith("abc", "b"))
	fmt.Println(StartsWith("abc", ""))
	fmt.Println(StartsWith("", ""))
	fmt.Println(StartsWith("", "a"))

	// Output:
	// true
	// false
	// true
	// true
	// false
}
