package strings

import "fmt"

func ExampleContainsChain() {
	fmt.Println(ContainsChain("abaabab", "a", 0))
	fmt.Println(ContainsChain("abaabab", "a", 1))
	fmt.Println(ContainsChain("abaabab", "a", 2))
	fmt.Println(ContainsChain("abaabab", "a", 3))
	fmt.Println(ContainsChain("", "a", 0))
	fmt.Println(ContainsChain("", "a", 1))
	fmt.Println(ContainsChain("", "", 1))
	fmt.Println(ContainsChain("", "", 2))

	// Output:
	// true
	// true
	// true
	// false
	// true
	// false
	// true
	// true
}
