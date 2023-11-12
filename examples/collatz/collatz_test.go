package collatz

import "fmt"

// ExampleCollatz berechnet die Collatz-Folgen für n = 1..10.
func ExampleCollatz() {
	for n := 1; n <= 10; n++ {
		fmt.Printf("n == %d: %v\n", n, Collatz(n))
	}

	// Output:
	// n == 1: [1]
	// n == 2: [2 1]
	// n == 3: [3 10 5 16 8 4 2 1]
	// n == 4: [4 2 1]
	// n == 5: [5 16 8 4 2 1]
	// n == 6: [6 3 10 5 16 8 4 2 1]
	// n == 7: [7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1]
	// n == 8: [8 4 2 1]
	// n == 9: [9 28 14 7 22 11 34 17 52 26 13 40 20 10 5 16 8 4 2 1]
	// n == 10: [10 5 16 8 4 2 1]
}
