package tourofgo

import "fmt"

// Practice with closures

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	previous := 0
	current := 0

	return func() int {
		if current == 0 {
			current = 1
		} else {
			new_current := current + previous
			previous, current = current, previous
			current = new_current
		}

		return current
	}
}

func PrintFibonacci(n int) {
	f := Fibonacci()

	for i := 0; i < n; i++ {
		if i == n-1 {
			fmt.Println(f())
		} else {
			fmt.Print(f(), ", ")
		}
	}
}
