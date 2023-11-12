package ackermann

import "fmt"

// ExampleAck berechnet die Werte der Ackermann-Funktion für n,m = 0..3.
func ExampleAck() {
	for m := 0; m <= 3; m++ {
		fmt.Printf("m == %d\n======\n", m)
		for n := 0; n <= 3; n++ {
			fmt.Printf("n == %d: %d\n", n, Ack(m, n))
		}
		fmt.Println()
	}

	// Output:
	// m == 0
	// ======
	// n == 0: 1
	// n == 1: 2
	// n == 2: 3
	// n == 3: 4
	//
	// m == 1
	// ======
	// n == 0: 2
	// n == 1: 3
	// n == 2: 4
	// n == 3: 5
	//
	// m == 2
	// ======
	// n == 0: 3
	// n == 1: 5
	// n == 2: 7
	// n == 3: 9
	//
	// m == 3
	// ======
	// n == 0: 5
	// n == 1: 13
	// n == 2: 29
	// n == 3: 61
}

// ExampleAckTable berechnet die Ackermann-Tabelle für m,n = 0..3.
func ExampleAckTable() {
	table := AckTable(3)

	for _, row := range table {
		fmt.Println(row)
	}

	// Output:
	// [1 2 3 4]
	// [2 3 4 5]
	// [3 5 7 9]
	// [5 13 29 61]
}
