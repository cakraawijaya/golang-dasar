package main

import "fmt"

func greeting(nama string) {
	fmt.Println("Halo", nama, "selamat belajar Golang!")
}

func main() {
	// Pemanggilan fungsi
	greeting("Andi")
	greeting("Sari")
	greeting("Budi")
}