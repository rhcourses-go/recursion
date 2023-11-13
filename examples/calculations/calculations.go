package calculations

/* Die Funktionen in diesem Paket dienen der Demonstration von
   Rekursion anhand einfacher Beispiele.
   Sie sind nicht für den produktiven Einsatz gedacht.

   Bei jeder Funktion ist angegeben, welche Gleichungen die
   Korrektheit der Funktion zeigen. Diese Gleichungen sind die
   Grundlage für die Implementierung der Funktion.
*/

// Add1 berechnet die Summe von x und y.
// begin:Add1
func Add1(x, y int) int {

	// Gleichungen für die Addition:
	// x + 0 = x
	// x + (y+1) = (x+y) + 1

	if y == 0 {
		return x
	}
	return Add1(x, y-1) + 1
}

// end:Add1

// Add2 berechnet die Summe von x und y.
// begin:Add2
func Add2(x, y int) int {

	// Gleichungen für die Addition:
	// x + 0 = x
	// x + y = (x+1) + (y-1)

	if y == 0 {
		return x
	}
	return Add2(x+1, y-1)
}

// end:Add2

// Power berechnet x hoch n.
// begin:Power
func Power(x, n int) int {

	// Gleichungen für die Potenz:
	// x^0 = 1
	// x^(n+1) = x^n * x

	if n == 0 {
		return 1
	}
	return Power(x, n-1) * x
}

// end:Power
