package main

import "fmt"

func main() {
	// Membuat array dengan 6 elemen
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:]		// ini untuk mengambil seluruh elemen
	fmt.Println("Ini adalah S1 :", s1)

	s2 := arr[:3]		// ini untuk mengambil elemen mulai dari index-0 hingga index-2
	fmt.Println("Ini adalah S2 :", s2)

	s3 := arr[2:]		// ini untuk mengambil elemen mulai dari index-2 hingga index akhir
	fmt.Println("Ini adalah S3 :", s3)

	s4 := arr[1:4]	// ini untuk mengambil elemen mulai dari index-1 hingga index-3
	fmt.Println("Ini adalah S4 :", s4)
}