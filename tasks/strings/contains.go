package strings

// Contains prüft, ob der String s die Sequenz seq enthält.
func Contains(s, seq string) bool {
	// begin:hint
	// Wenn seq leer ist, ist das Ergebnis true.
	// Wenn s leer ist, ist das Ergebnis false.
	// Wenn s mit seq beginnt, ist das Ergebnis true.
	//   (Verwenden Sie die Hilfsfunktion StartsWith aus startswith.go.)
	// Wenn s nicht mit seq beginnt, ist das Ergebnis Contains(s[1:], seq).
	// end:hint
	// begin:solution
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	return StartsWith(s, seq) || Contains(s[1:], seq)
	// end:solution
	// iftask: return true
}
