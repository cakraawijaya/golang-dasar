package main

import "fmt"

func main() {
	// User mempresentasikan data pengguna
	type User struct {
		Name	string
		Email	string
		Age		int
	}

	// Membuat object User
	user1 := User {
		Name:  "Andi",
		Email: "andi@gmail.com",
		Age:   25,
	}

	// Menampilkan data user
	fmt.Println("Nama  :", user1.Name)
	fmt.Println("Email :", user1.Email)
	fmt.Println("Umur  :", user1.Age)
}