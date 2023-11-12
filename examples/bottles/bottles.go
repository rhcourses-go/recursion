package bottles

import "fmt"

// Bottles gibt den Text des Liedes "99 Bottles of Beer" aus.
// Der Text startet bei der angegebenen Anzahl von Flaschen.
func Bottles(n int) {
	if n < 0 {
		return
	}

	// Rekursionsanker: n == 0
	// Abschluss des Liedes, wenn keine Flaschen mehr übrig sind.
	if n == 0 {
		fmt.Println("No more bottles of beer on the wall, no more bottles of beer.")
		fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall.")
		return
	}

	// Rekursionsschritt
	// Spezialfall, wenn nur noch eine Flasche übrig ist.
	// In diesem Fall wird "bottle" statt "bottles" verwendet.
	if n == 1 {
		fmt.Println("1 bottle of beer on the wall, 1 bottle of beer.")
		fmt.Println("Take one down and pass it around, no more bottles of beer on the wall.")
		fmt.Println()
		Bottles(n - 1)
		return
	}

	// Rekursionsschritt
	// Allgemeiner Fall, wenn noch mehr als eine Flasche übrig sind.
	fmt.Printf("%d bottles of beer on the wall, %d bottles of beer.\n", n, n)
	fmt.Printf("Take one down and pass it around, %d bottles of beer on the wall.\n", n-1)
	fmt.Println()
	Bottles(n - 1)
}
