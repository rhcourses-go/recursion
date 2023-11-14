package strings

// StartsWith liefert true, falls der string s mit der Sequenz seq beginnt.
func StartsWith(s, seq string) bool {
	// begin:hint
	// Wenn seq leer ist, ist das Ergebnis true.
	// Wenn s leer ist, ist das Ergebnis false.
	// Wenn s[0] != seq[0], ist das Ergebnis false.
	// Wenn s[0] == seq[0], muss geprüft werden, ob s[1:] mit seq[1:] beginnt.
	// end:hint
	// begin:solution
	if seq == "" {
		return true
	}
	if s == "" {
		return false
	}
	return s[0] == seq[0] && StartsWith(s[1:], seq[1:])
	// end:solution
	// iftask: return true
}
