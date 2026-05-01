package main

import "fmt"

func main() {
	x := 3
	fmt.Println("Nilai x =", x)

	if x > 5 {
		fmt.Println("x lebih besar dari 5")
	} else {
		fmt.Println("x tidak lebih besar dari 5")
	}
}