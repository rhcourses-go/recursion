package lists

// Liefert true, falls die beiden Listen gleich sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func ListsEqual(list1, list2 []int) bool {
	// begin:hint
	// Wenn beide Listen leer sind, sind sie gleich.
	// Wenn eine Liste leer ist und die andere nicht, sind sie nicht gleich.
	// Wenn beide Listen nicht leer sind, müssen wir prüfen,
	// ob die ersten Elemente gleich sind und ob die Restlisten gleich sind.
	// end:hint
	// begin:solution
	if Empty(list1) && Empty(list2) {
		return true
	}
	if Empty(list1) && !Empty(list2) || !Empty(list1) && Empty(list2) {
		return false
	}
	if list1[0] != list2[0] {
		return false
	}
	return ListsEqual(list1[1:], list2[1:])
	// end:solution
	// iftask: return true
}
