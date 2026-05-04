package main

import "fmt"

func main() {
	angka := 10
	fmt.Println("Input =", angka)

	switch angka {
		case 1:
			fmt.Println("Angka 1")
		case 2:
			fmt.Println("Angka 2")
		case 3:
			fmt.Println("Angka 3")
		case 4:
			fmt.Println("Angka 4")
		case 5:
			fmt.Println("Angka 5")
		default:
			fmt.Println("Angka tidak dikenali")
	}
}