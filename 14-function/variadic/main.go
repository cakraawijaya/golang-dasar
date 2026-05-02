package main

import "fmt"

// Fungsi variadic untuk menjumlahkan semua angka
func jumlahAngka(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v	// total = total + v
	}
	return total
}

func main() {
	fmt.Println(jumlahAngka(1, 2, 3)) // 6
	fmt.Println(jumlahAngka(10, 20, 30, 5)) // 65
	fmt.Println(jumlahAngka()) // 0
}