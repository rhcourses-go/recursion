package countdown

import "fmt"

// CountDown zählt rekursiv von n bis 0 herunter.
// begin:CountDown
func CountDown(n int) {
	if n <= 0 {
		return
	}
	fmt.Println(n)
	CountDown(n - 1)
}

// end:CountDown
