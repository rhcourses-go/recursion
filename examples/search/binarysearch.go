package search

// FindSorted sucht in einer sortierten Liste nach dem ersten Vorkommen von x.
// Falls x nicht gefunden wird, wird -1 zurückgegeben.
// Da die Liste sortiert ist, wird die binäre Suche verwendet.
func FindSorted(list []int, x int) int {

	// Rekursionsanker: Abbruch falls Liste leer.
	if len(list) == 0 {
		return -1
	}

	// Rekursionsanker: Abbruch falls das mittlere Element x ist.
	c := len(list) / 2
	if list[c] == x {
		return c
	}

	// Rekursionsschritt: Suche in der linken oder rechten Hälfte.
	sublist := list[:c]
	offset := 0 // Offset für den Index in der Teilliste
	if x > list[c] {
		sublist = list[c+1:]
		offset = c + 1
	}
	pos := FindSorted(sublist, x)

	// Wurde x nicht gefunden, wird -1 zurückgegeben.
	if pos == -1 {
		return -1
	}

	// Falls x gefunden wurde, muss der Index um den Offset erhöht werden.
	return pos + offset
}
