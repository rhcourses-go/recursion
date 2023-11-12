package factorial

// Factorial berechnet die Fakultät von n.
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
