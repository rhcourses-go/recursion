package hanoi

import "fmt"

// Move führt eine Scheibenbewegung von s nach z aus.
// In diesem Fall bedeutet das, dass die Bewegung auf der Konsole ausgegeben wird.
// begin:Move
func Move(s, z string) {
	fmt.Printf("Bewege Scheibe von %s nach %s.\n", s, z)
}

// end:Move

/* Die folgenden Funktionen lösen das Problem jeweils für eine bestimmte Höhe.
Jede Funktion erwartet als Parameter die Name der drei Türme in der
Reihenfolge "start", "mitte", "ziel" (s,m,z).
*/

// Hanoi1 löst das Türme-von-Hanoi-Problem für Höhe 1.
// begin:Hanoi1
func Hanoi1(s, m, z string) {
	Move(s, z)
}

// end:Hanoi1

// Hanoi2 löst das Türme-von-Hanoi-Problem für Höhe 2.
// begin:Hanoi2
func Hanoi2(s, m, z string) {
	Hanoi1(s, z, m)
	Move(s, z)
	Hanoi1(m, s, z)
}

// end:Hanoi2

// Hanoi3 löst das Türme-von-Hanoi-Problem für Höhe 3.
// begin:Hanoi3
func Hanoi3(s, m, z string) {
	Hanoi2(s, z, m)
	Move(s, z)
	Hanoi2(m, s, z)
}

// end:Hanoi3

// Hanoi4 löst das Türme-von-Hanoi-Problem für Höhe 4.
// begin:Hanoi4
func Hanoi4(s, m, z string) {
	Hanoi3(s, z, m)
	Move(s, z)
	Hanoi3(m, s, z)
}

// end:Hanoi4

// Hanoi5 löst das Türme-von-Hanoi-Problem für Höhe 5.
// begin:Hanoi5
func Hanoi5(s, m, z string) {
	Hanoi4(s, z, m)
	Move(s, z)
	Hanoi4(m, s, z)
}

// end:Hanoi5

// Hanoi6 löst das Türme-von-Hanoi-Problem für Höhe 6.
// begin:Hanoi6
func Hanoi6(s, m, z string) {
	Hanoi5(s, z, m)
	Move(s, z)
	Hanoi5(m, s, z)
}

// end:Hanoi6

/* Die folgende Funktion löst das Problem für eine beliebige Höhe.
   Sie erwartet als Parameter die Höhe und die Nummern der drei Türme in der
   Reihenfolge "von", "über", "nach".
   Es fällt auf, dass die Funktion sich selbst aufruft, ansonsten aber
   das gleiche Muster wie die Funktionen für die einzelnen Höhen aufweist.
*/

// Hanoi löst das Türme-von-Hanoi-Problem für eine beliebige Höhe h.
// begin:HanoiComplete
func Hanoi(h int, s, m, z string) {
	if h == 1 {
		Move(s, z)
	} else {
		Hanoi(h-1, s, z, m)
		Move(s, z)
		Hanoi(h-1, m, s, z)
	}
}

// end:HanoiComplete
