package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai ujian: ")
	fmt.Scanln(&nilai)

	if v := nilai; v >= 90 {
		fmt.Println("Nilai Anda: A (Sangat Baik)")
	} else if v >= 75 {
		fmt.Println("Nilai Anda: B (Baik)")
	} else if v >= 60 {
		fmt.Println("Nilai Anda: C (Cukup)")
	} else {
		fmt.Println("Nilai Anda: D (Perlu Perbaikan)")
	}
}