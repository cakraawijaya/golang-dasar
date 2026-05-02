package main

import "fmt"

func tanganiPanic() {
	if r := recover(); r != nil {
		fmt.Println("Hasil: Terjadi panic, tapi sudah ditangani.", r)
	}
}

func bagi(a, b int) {
	defer tanganiPanic() // defer akan memastikan tanganiPanic dipanggil walaupun terjadi panic
	fmt.Printf("Membagi %d dengan %d\n", a, b)
	hasil := a / b // ini bisa menyebabkan panic kalau b = 0
	fmt.Println("Hasil:", hasil)
}

func main() {
	fmt.Println("Program mulai")
	bagi(10, 2) // normal
	bagi(5, 0) // ini akan menyebabkan panic, tapi recover akan menangkapnya
	fmt.Println("Program selesai dengan aman")
}