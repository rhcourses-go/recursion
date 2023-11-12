package search

// FindMin liefert das Minimum einer Liste.
// Falls die Liste leer ist, wird -1 zurückgegeben.
func FindMin(list []int) int {
	if len(list) == 0 {
		return -1
	}

	head, tail := list[0], list[1:]
	min := FindMin(tail)

	if min == -1 || head < min {
		return head
	}
	return min
}
