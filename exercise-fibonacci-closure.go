package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	old_value := 0
	next_value := 1
	return func() int {
		latest_value := old_value + next_value
		old_value, next_value = next_value, latest_value
		return latest_value
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
