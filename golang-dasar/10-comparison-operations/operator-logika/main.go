package main

import "fmt"

func main() {
	var angka1, angka2 int

	// Input dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)

	fmt.Println("\n=== Hasil logika ===")

	// AND (&&)
	fmt.Printf("(%d > 5) && (%d < 10) ? %v\n", angka1, angka2, angka1 > 5 && angka2 < 10)

	// OR (||)
	fmt.Printf("(%d > 5) || (%d < 10) ? %v\n", angka1, angka2, angka1 > 5 || angka2 < 10)

	// NOT (!)
	fmt.Printf("!(%d > %d) ? %v\n", angka1, angka2, !(angka1 > angka2))
}