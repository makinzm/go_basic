package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

// https://go-tour-jp.appspot.com/methods/20
func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return fmt.Sprintf("cannot Sqrt negative number: %f", e)
	} else {
		return ""
	}
}

func Sqrt(x float64) (float64, error) {
	err := ErrNegativeSqrt(x)
	return x, err
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
