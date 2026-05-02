package main

import "fmt"

func main() {
	// membuat variabel baru bernama x dengan nilai 10 (tipe data: integer)
	x := 10

	x = 5	// mengubah nilai x dari 10 menjadi 5
	fmt.Println("Nilai X  = ", x)

	x += 2	// x = x + 2 -> x = 5 + 2 = 7
	fmt.Println("\nx += 2 --> ", x)

	x -= 2	// x = x - 2 -> x = 7 - 2 = 5
	fmt.Println("x -= 2 --> ", x)

	x *= 2	// x = x * 2 -> x = 5 * 2 = 10
	fmt.Println("x *= 2 --> ", x)

	x /= 2	// x = x / 2 -> x = 10 / 2 = 5
	fmt.Println("x /= 2 --> ", x)

	x %= 2	// x = x % 2 -> x = 5 % 2 = 1
	fmt.Println("x %= 2 --> ", x)
}