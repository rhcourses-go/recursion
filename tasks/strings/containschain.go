package strings

// ContainsChain liefert true, falls s eine Kette von count aufeinanderfolgenden
// Vorkommen von symbol enthält.
func ContainsChain(s, symbol string, count int) bool {
	// begin:hint
	// Wenn count == 0, ist das Ergebnis true.
	// Wenn s leer ist, ist das Ergebnis false.
	//
	// Wenn s nicht leer ist, gibt es zwei Möglichkeiten, wie true geliefert werden kann:
	// 1. s beginnt mit symbol und enthält danach eine Kette von
	//    count-1 aufeinanderfolgenden Vorkommen von symbol.
	// 2. s beginnt nicht mit symbol, aber s[1:] enthält eine Kette von
	//    count aufeinanderfolgenden Vorkommen von symbol.
	// end:hint
	// begin:solution
	if count == 0 {
		return true
	}
	if s == "" {
		return symbol == ""
	}
	return (s[:1] == symbol && StartsWith(s[:1], Chain(symbol, count-1))) || ContainsChain(s[1:], symbol, count)
	// end:solution
	// iftask: return true
}
