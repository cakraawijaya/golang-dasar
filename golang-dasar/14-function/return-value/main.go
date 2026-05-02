package main

import "fmt"

func tambah(a, b int) (hasil int) {
	hasil = a + b
	// cukup tulis return saja
	return
}

func main() {
	fmt.Println(tambah(3, 5)) // 8
}