package ackermann

// Ack berechnet die Ackermann-Funktion für m und n.
func Ack(m, n int) int {
	if m == 0 {
		return n + 1
	}
	if n == 0 {
		return Ack(m-1, 1)
	}
	return Ack(m-1, Ack(m, n-1))
}

// AckTable berechnet die Werte der Ackermann-Funktion für m,n = 0..max.
func AckTable(max int) [][]int {
	table := make([][]int, max+1)
	for m := 0; m <= max; m++ {
		table[m] = make([]int, max+1)
		for n := 0; n <= max; n++ {
			table[m][n] = Ack(m, n)
		}
	}
	return table
}
