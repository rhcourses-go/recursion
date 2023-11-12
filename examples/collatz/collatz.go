package collatz

// Collatz berechnet die Collatz-Folge für n.
func Collatz(n int) []int {
	// Spezialfall: Kein Ergebnis für n <= 0
	if n <= 0 {
		return nil
	}

	// Rekursionsanker: n == 1
	if n == 1 {
		return []int{1}
	}

	// Rekursionsschritt, falls n gerade
	if n%2 == 0 {
		return append([]int{n}, Collatz(n/2)...)
	}

	// Rekursionsschritt, falls n ungerade
	return append([]int{n}, Collatz(3*n+1)...)
}
