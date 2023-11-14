package lists

// Liefert eine Liste, die alle Elemente aus list enthält,
// außer dem an Stelle pos.
// Wenn pos außerhalb des Bereichs der Liste liegt, wird die
// ursprüngliche Liste zurückgegeben.
// Verwenden Sie Rekursion und benutzen Sie NICHT die len-Funktion.
// Sie können die Hilfsfunktion Empty aus empty.go verwenden.
func RemoveElement(list []int, pos int) []int {
	// begin:hint
	// Wenn die Liste leer ist, ist das Ergebnis die leere Liste.
	// Wenn pos 0 ist, ist das Ergebnis die Restliste.
	// Wenn pos größer als 0 ist, ist das Ergebnis das erste Element
	// plus die Restliste ohne das Element an Stelle pos-1.
	// end:hint
	// begin:solution
	if Empty(list) {
		return list
	}
	if pos == 0 {
		return list[1:]
	}
	return append(list[:1], RemoveElement(list[1:], pos-1)...)
	// end:solution
	// iftask: return list
}
