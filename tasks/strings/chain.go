package strings

// Chain erwartet einen String und hängt ihn n mal hintereinander.
// Liefert das Ergebnis.
func Chain(s string, n int) string {
	// begin:hint
	// Wenn n == 0, ist das Ergebnis der leere String.
	// Wenn n > 0, ist das Ergebnis der String plus Chain(s, n-1).
	// end:hint
	// begin:solution
	if n == 0 {
		return ""
	}
	return s + Chain(s, n-1)
	// end:solution
	// iftask: return ""
}
