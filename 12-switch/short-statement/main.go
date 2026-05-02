package main

import "fmt"

func main() {
	switch panjang := len("Geologi"); {
		case panjang > 10:
			fmt.Println("Kata terlalu panjang")
		case panjang >= 5:
			fmt.Println("Kata sedang")
		default:
			fmt.Println("Kata pendek")
	}
}