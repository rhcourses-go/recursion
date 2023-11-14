package search

// Find sucht in einer Liste nach dem ersten Vorkommen von x
// und gibt dessen Index zurück. Falls x nicht gefunden wird,
// wird -1 zurückgegeben.
func Find(list []int, x int) int {

	// begin:hint
	// Rekursionsanker: Abbruch falls Liste leer.
	// Rekursionsanker: Abbruch falls das erste Element x ist.
	// Rekursionsschritt: Suche in der Restliste.
	// end:hint
	// begin:solution
	if len(list) == 0 {
		return -1
	}
	if list[0] == x {
		return 0
	}
	pos := Find(list[1:], x)

	if pos == -1 {
		return -1
	}
	return Find(list[1:], x) + 1
	// end:solution
	// iftask: return -1
}
