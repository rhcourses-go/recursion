package factorial

import "fmt"

// ExampleFactorial berechnet die Fakultäten von 0 bis 10 und gibt sie aus.
func ExampleFactorial() {
	for i := 0; i <= 10; i++ {
		fmt.Println(Factorial(i))
	}
	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
	// 720
	// 5040
	// 40320
	// 362880
	// 3628800
}
