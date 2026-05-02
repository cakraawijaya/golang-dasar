package main

import "fmt"

func main() {
	x := 1 // x = 1
	defer fmt.Println("defer-1, x =", x)	// argumen x dievaluasi sekarang (1)
	x = 2 // x = 2
	defer fmt.Println("defer-2, x =", x)	// argumen x dievaluasi sekarang (2)
	fmt.Println("body, x =", x)
}