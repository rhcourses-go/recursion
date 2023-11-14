package calc

// BinomialCoefficient erwartet zwei Zahlen n und k und liefert den
// Binomialkoeffizienten "n über k".
func BinomialCoefficient(n, k int) int {
	// begin:hint
	// Berechnen Sie den Binomialkoeffizienten rekursiv.
	// Es gilt: (n,k) == (n-1,k-1) + (n-1,k)
	// Diese Berechnungsvorschrift können Sie direkt in eine Rekursion umsetzen,
	// Sie müssen nur noch den Rekursionsanfang hinzufügen.
	// Anmerkung: Diese Berechnungsvorschrift beschreibt exakt das Pascalsche Dreieck!
	// end:hint
	// begin:solution
	if n == 0 {
		return 1
	}
	if k == 0 || k == n {
		return 1
	}
	return BinomialCoefficient(n-1, k-1) + BinomialCoefficient(n-1, k)
	// end:solution
	// iftask: return 0
}
