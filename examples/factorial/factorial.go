package factorial

// Factorial berechnet die Fakultät von n.
// begin:Factorial
func Factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}

// end:Factorial

// FactorialIter berechnet die Fakultät von n iterativ.
// begin:FactorialIter
func FactorialIter(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// end:FactorialIter
