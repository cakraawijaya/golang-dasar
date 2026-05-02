package main

import (
	"fmt"
	"math"
)

func main() {
	a := 10
	b := 3

	fmt.Println("Operasi Dasar Matematika dengan tipe data Integer:")
	fmt.Printf("a + b = %d + %d = %d\n", a, b, a+b)	// Penjumlahan
	fmt.Printf("a - b = %d - %d = %d\n", a, b, a-b)	// Pengurangan
	fmt.Printf("a x b = %d x %d = %d\n", a, b, a*b)	// Perkalian
	fmt.Printf("a / b = %d / %d = %d\n", a, b, a/b)	// Pembagian
	fmt.Printf("a %% b = %d %% %d = %d\n", a, b, a%b)	// Modulus

	// Dengan float
	x := 10.0
	y := 3.0

	fmt.Println("\nOperasi Dasar Matematika dengan tipe data Float:")
	fmt.Printf("x + y = %.1f + %.1f = %.1f\n", x, y, x+y)	// Penjumlahan
	fmt.Printf("x - y = %.1f - %.1f = %.1f\n", x, y, x-y)	// Pengurangan
	fmt.Printf("x x y = %.1f x %.1f = %.1f\n", x, y, x*y)	// Perkalian
	fmt.Printf("x / y = %.1f / %.1f = %.1f\n", x, y, x/y)	// Pembagian
	fmt.Printf("x %% y = %.1f %% %.1f = %.1f\n", x, y, math.Mod(x, y))	// Modulus
}