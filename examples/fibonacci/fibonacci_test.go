package fibonacci

import "fmt"

// ExampleFib berechnet die ersten 10 Fibonacci-Zahlen
// mittels der Funktion Fib und gibt sie aus.
func ExampleFib() {
	for i := 0; i < 10; i++ {
		fmt.Println(Fib(i))
	}
	// Output:
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// 13
	// 21
	// 34
	// 55
}

// ExampleFibNumbers berechnet die Liste der ersten 10 Fibonacci-Zahlen
// mittels der Funktion FibNumbers und gibt sie aus.
func ExampleFibNumbers() {
	fmt.Println(FibNumbers(10))
	// Output: [1 1 2 3 5 8 13 21 34 55]
}

// ExampleFibRatios berechnet die Liste der ersten 10 Fibonacci-Verhältnisse
// mittels der Funktion FibRatios und gibt sie mit 5 Stellen Genauigkeit aus.
func ExampleFibRatios() {
	ratios := FibRatios(11) // bis 11, damit es 10 Verhältnisse gibt
	for _, r := range ratios {
		fmt.Printf("%.5f\n", r)
	}

	// Output:
	// 1.00000
	// 2.00000
	// 1.50000
	// 1.66667
	// 1.60000
	// 1.62500
	// 1.61538
	// 1.61905
	// 1.61765
	// 1.61818
}
