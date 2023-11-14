package calc

// Liefert das Produkt von x und y.
func Mult(x, y int) int {
	result := 1
	// begin:hint
	// x * 0 = 0
	// x * (y + 1) = (x * y) + x
	// end:hint
	// begin:solution
	if y == 0 {
		return 0
	}
	result = Mult(x, y-1) + x
	// end:solution
	return result
}
