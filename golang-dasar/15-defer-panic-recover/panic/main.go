package main

import "fmt"

func main() {
	fmt.Println("defer ini tetap jalan sebelum program mati")
	panic("ada sesuatu yang fatal")
	// baris di bawah tidak dieksekusi
	fmt.Println("Setelah panic")
}