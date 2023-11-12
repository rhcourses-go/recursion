package hanoi

import "fmt"

// Move führt eine Scheibenbewegung aus.
// In diesem Fall bedeutet das, dass die Bewegung auf der Konsole ausgegeben wird.
func Move(from, to string) {
	fmt.Printf("Bewege Scheibe von %s nach %s.\n", from, to)
}

/* Die folgenden Funktionen lösen das Problem jeweils für eine bestimmte Höhe.
   Jede Funktion erwartet als Parameter die Nummern der drei Türme in der
   Reihenfolge "von", "über", "nach".
*/

// Hanoi1 löst das Türme-von-Hanoi-Problem für Höhe 1.
func Hanoi1(from, via, to string) {
	Move(from, to)
}

// Hanoi2 löst das Türme-von-Hanoi-Problem für Höhe 2.
func Hanoi2(from, via, to string) {
	Hanoi1(from, to, via)
	Move(from, to)
	Hanoi1(via, from, to)
}

// Hanoi3 löst das Türme-von-Hanoi-Problem für Höhe 3.
func Hanoi3(from, via, to string) {
	Hanoi2(from, to, via)
	Move(from, to)
	Hanoi2(via, from, to)
}

// Hanoi4 löst das Türme-von-Hanoi-Problem für Höhe 4.
func Hanoi4(from, via, to string) {
	Hanoi3(from, to, via)
	Move(from, to)
	Hanoi3(via, from, to)
}

// Hanoi5 löst das Türme-von-Hanoi-Problem für Höhe 5.
func Hanoi5(from, via, to string) {
	Hanoi4(from, to, via)
	Move(from, to)
	Hanoi4(via, from, to)
}

// Hanoi6 löst das Türme-von-Hanoi-Problem für Höhe 6.
func Hanoi6(from, via, to string) {
	Hanoi5(from, to, via)
	Move(from, to)
	Hanoi5(via, from, to)
}

/* Die folgende Funktion löst das Problem für eine beliebige Höhe.
   Sie erwartet als Parameter die Höhe und die Nummern der drei Türme in der
   Reihenfolge "von", "über", "nach".
   Es fällt auf, dass die Funktion sich selbst aufruft, ansonsten aber
   das gleiche Muster wie die Funktionen für die einzelnen Höhen aufweist.
*/

// Hanoi löst das Türme-von-Hanoi-Problem für eine beliebige Höhe.
func Hanoi(height int, from, via, to string) {
	if height == 1 {
		Move(from, to)
	} else {
		Hanoi(height-1, from, to, via)
		Move(from, to)
		Hanoi(height-1, via, from, to)
	}
}
