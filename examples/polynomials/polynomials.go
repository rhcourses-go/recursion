package polynomials

// Polynomial ist ein Polynom.
// Es ist als Liste seiner Koeffizienten gegeben.
// Der Koeffizient des höchsten Terms steht an erster Stelle.
type Polynomial []float64

// Evaluate berechnet den Wert des Polynoms p an der Stelle x.
func Evaluate(p Polynomial, x float64) float64 {
	if len(p) == 0 {
		return 0
	}
	last := len(p) - 1
	return p[last] + x*Evaluate(p[:last], x)
}
