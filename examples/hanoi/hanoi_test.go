package hanoi

import "fmt"

// ExampleMove zeigt die Funktionsweise von Move.
func ExampleMove() {
	Move("A", "B")
	Move("A", "C")
	Move("B", "C")

	// Output:
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
}

// ExampleHanoi1 zeigt die Funktionsweise von Hanoi1.
func ExampleHanoi1() {
	Hanoi1("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach C.
}

// ExampleHanoi2 zeigt die Funktionsweise von Hanoi2.
func ExampleHanoi2() {
	Hanoi2("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
}

// ExampleHanoi3 zeigt die Funktionsweise von Hanoi3.
func ExampleHanoi3() {
	Hanoi3("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
}

// ExampleHanoi4 zeigt die Funktionsweise von Hanoi4.
func ExampleHanoi4() {
	Hanoi4("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
}

// ExampleHanoi5 zeigt die Funktionsweise von Hanoi5.
func ExampleHanoi5() {
	Hanoi5("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
}

// ExampleHanoi6 zeigt die Funktionsweise von Hanoi6.
func ExampleHanoi6() {
	Hanoi6("A", "B", "C")

	// Output:
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von C nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
}

// ExampleHanoi zeigt die Funktionsweise von Hanoi.
func ExampleHanoi() {
	fmt.Println("Hanoi1:")
	fmt.Println("=======")
	Hanoi(1, "A", "B", "C")

	fmt.Println()
	fmt.Println("Hanoi2:")
	fmt.Println("=======")
	Hanoi(2, "A", "B", "C")

	fmt.Println()
	fmt.Println("Hanoi3:")
	fmt.Println("=======")
	Hanoi(3, "A", "B", "C")

	// Output:
	// Hanoi1:
	// =======
	// Bewege Scheibe von A nach C.
	//
	// Hanoi2:
	// =======
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach C.
	//
	// Hanoi3:
	// =======
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von A nach B.
	// Bewege Scheibe von C nach B.
	// Bewege Scheibe von A nach C.
	// Bewege Scheibe von B nach A.
	// Bewege Scheibe von B nach C.
	// Bewege Scheibe von A nach C.
}
