package permutations

import (
	"fmt"
	"slices"
	"strings"
)

// ExamplePermutations zeigt die Funktionsweise von Permutations.
func ExamplePermutations() {
	s1 := "abcd"
	p1 := Permutations(s1)
	slices.Sort(p1)
	fmt.Printf("Ausgangs-String: %s\n", s1)
	fmt.Println(strings.Repeat("=", len(s1)+17))
	for _, s := range p1 {
		fmt.Println(s)
	}

	fmt.Println()

	s2 := "abc"
	p2 := Permutations(s2)
	slices.Sort(p2)
	fmt.Printf("Ausgangs-String: %s\n", s2)
	fmt.Println(strings.Repeat("=", len(s2)+17))
	for _, s := range p2 {
		fmt.Println(s)
	}

	// Output:
	// Ausgangs-String: abcd
	// =====================
	// abcd
	// abdc
	// acbd
	// acdb
	// adbc
	// adcb
	// bacd
	// badc
	// bcad
	// bcda
	// bdac
	// bdca
	// cabd
	// cadb
	// cbad
	// cbda
	// cdab
	// cdba
	// dabc
	// dacb
	// dbac
	// dbca
	// dcab
	// dcba
	//
	// Ausgangs-String: abc
	// ====================
	// abc
	// acb
	// bac
	// bca
	// cab
	// cba
}
