package lists

// Liefert eine Liste mit allen Elementen aus list, die kleiner oder gleich key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterLess(list []int, key int) []int {
	// begin:hint
	// Gehen Sie ähnlich wie bei Remove vor:
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.
	// Wenn das erste Element größer als key ist, ist das Ergebnis die gefilterte Restliste.
	// Wenn das erste Element kleiner oder gleich key ist, ist das Ergebnis das erste Element
	// plus die gefilterte Restliste.
	// end:hint
	// begin:solution
	if Empty(list) {
		return list
	}
	if list[0] > key {
		return FilterLess(list[1:], key)
	}
	return append(list[:1], FilterLess(list[1:], key)...)
	// end:solution
	// iftask: return list
}

// Liefert eine Liste mit allen Elementen aus list, die echt größer als key sind.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func FilterGreater(list []int, key int) []int {
	// begin:hint
	// Gehen Sie analog zu FilterLess vor.
	// end:hint
	// begin:solution
	if Empty(list) {
		return list
	}
	if list[0] <= key {
		return FilterGreater(list[1:], key)
	}
	return append(list[:1], FilterGreater(list[1:], key)...)
	// end:solution
	// iftask: return list
}
