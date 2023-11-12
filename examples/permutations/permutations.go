package permutations

// Permutations erwartet einen String und berechnet alle Permutationen dieser Zeichen.
func Permutations(elements string) []string {

	// Rekursionsanker: Ein String mit einem oder keinem Zeichen
	// hat nur eine Permutation. Wichtig: Auch bei Länge 0 muss
	// ein Slice zurückgegeben werden, damit die Rekursion etwas
	// anzuhängen hat.
	if len(elements) <= 1 {
		return []string{elements}
	}

	// Rekursionsschritt: Für jedes Zeichen des Strings wird ein
	// neuer String berechnet, in dem das Zeichen fehlt. Die
	// Permutationen dieser neuen Strings werden mit dem Zeichen
	// kombiniert und zurückgegeben.
	permutations := []string{}
	for i, element := range elements {

		// Kopie des Strings ohne das aktuelle Zeichen
		shortString := elements[:i] + elements[i+1:]

		// Permutationen von shortString berechnen und
		// mit dem aktuellen Zeichen kombinieren
		for _, permutation := range Permutations(shortString) {
			permutations = append(permutations, string(element)+permutation)
		}
	}
	return permutations
}
