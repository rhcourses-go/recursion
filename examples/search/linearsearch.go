package search

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {

	// Rekursionsanker: Abbruch falls Liste leer.
	if len(list) == 0 {
		return -1
	}

	// Rekursionsanker: Abbruch falls das erste Element x ist.
	if list[0] == x {
		return 0
	}

	// Rekursionsschritt: Suche in der Restliste.
	pos := Find(list[1:], x)

	// Wurde x nicht gefunden, wird -1 zurückgegeben.
	if pos == -1 {
		return -1
	}

	// Falls x gefunden wurde, muss der Index um 1 erhöht werden,
	// da das erste Element der Restliste dem zweiten Element der
	// ursprünglichen Liste entspricht.
	return Find(list[1:], x) + 1
}

// FindTail ist eine alternative Implementierung von Find.
// Diese Art der Implementierung wird auch als "Tail Recursion"
// bezeichnet, da der rekursive Aufruf am Ende der Funktion
// steht. Der Compiler kann diese Art der Rekursion optimieren,
// so dass kein zusätzlicher Stack-Frame benötigt wird.
// Allerdings ist diese Art der Rekursion nicht so einfach zu
// lesen und ähnelt sehr stark einer Iteration.
func FindTail(list []int, x int) int {
	return findTail(list, x, 0)
}

// findTail ist die eigentliche Implementierung von FindTail.
// Der Parameter pos enthält den aktuellen Index in der Liste.
func findTail(list []int, x int, pos int) int {

	if len(list) == 0 || pos >= len(list) {
		return -1
	}

	if list[pos] == x {
		return pos
	}

	return findTail(list, x, pos+1)
}
