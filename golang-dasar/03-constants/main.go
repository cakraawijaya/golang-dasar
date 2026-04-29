package main

import "fmt"

func main() {
	const pi = 3.14
	const appName = "Project Manager App"

	fmt.Println("Nilai Pi:", pi)
	fmt.Println("Nama Aplikasi:", appName)

	// Jika coba diubah -> ERROR
	// pi = 3.1415 // ❌ tidak boleh
}