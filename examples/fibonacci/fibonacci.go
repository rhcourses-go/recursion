package fibonacci

// Fib berechnet die n-te  Fibonacci-Zahl.
func Fib(n int) int {
	if n <= 1 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// FibNumbers berechnet die ersten n Fibonacci-Zahlen.
func FibNumbers(n int) []int {
	fibs := make([]int, n)
	if n >= 1 {
		fibs[0] = 1
	}
	if n >= 2 {
		fibs[1] = 1
	}
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}

// FibRatios berechnet alle Verhältnisse zweier aufeinanderfolgender
// Fibonacci-Zahlen bis zur n-ten Fibonacci-Zahl.
func FibRatios(n int) []float64 {
	fibs := FibNumbers(n)
	ratios := make([]float64, n-1)
	for i := 0; i < n-1; i++ {
		ratios[i] = float64(fibs[i+1]) / float64(fibs[i])
	}
	return ratios
}
