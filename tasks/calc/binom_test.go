package calc

import "fmt"

// ExampleBinomialCoefficient berechnet mit Hilfe der Funktion BinomialCoefficient
// die ersten Zeilen des Pascalschen Dreiecks und gibt sie aus.
func ExampleBinomialCoefficient() {
	for n := 0; n <= 5; n++ {
		fmt.Printf("n == %d:", n)
		for k := 0; k <= n; k++ {
			fmt.Printf(" %d", BinomialCoefficient(n, k))
		}
		fmt.Println()
	}

	// Output:
	// n == 0: 1
	// n == 1: 1 1
	// n == 2: 1 2 1
	// n == 3: 1 3 3 1
	// n == 4: 1 4 6 4 1
	// n == 5: 1 5 10 10 5 1
}
