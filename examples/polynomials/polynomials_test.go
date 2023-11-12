package polynomials

import "fmt"

// ExampleEvaluate zeigt die Verwendung von Evaluate.
func ExampleEvaluate() {
	p1 := Polynomial{3, 2, 1}       // 3x^2 + 2x + 1
	p2 := Polynomial{3, 0, 1}       // 3x^2 + 1
	p3 := Polynomial{1, 0, 0}       // x^2
	p4 := Polynomial{1, 1, 0, 0, 1} // x^4 + x^3 + 1

	fmt.Println(Evaluate(p1, 0))
	fmt.Println(Evaluate(p1, 5))
	fmt.Println()

	fmt.Println(Evaluate(p2, 0))
	fmt.Println(Evaluate(p2, 5))
	fmt.Println()

	fmt.Println(Evaluate(p3, 0))
	fmt.Println(Evaluate(p3, 5))
	fmt.Println()

	fmt.Println(Evaluate(p4, 0))
	fmt.Println(Evaluate(p4, 5))

	// Output:
	// 1
	// 86
	//
	// 1
	// 76
	//
	// 0
	// 25
	//
	// 1
	// 751
}
