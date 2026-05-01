package main

import "fmt"

func main() {
	var angka1, angka2 int

	// Input dari pengguna
	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&angka1)
	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&angka2)
	// Latihan perbandingan
	fmt.Println("\n=== Hasil perbandingan ===")
	fmt.Printf("%d == %d ? %v\n", angka1, angka2, angka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angka1, angka2, angka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angka1, angka2, angka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angka1, angka2, angka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angka1, angka2, angka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angka1, angka2, angka1 <= angka2)
}