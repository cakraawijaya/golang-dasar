package main

import "fmt"

func main() {
	type Address struct {
		City, Province string
	}

	// Struct User menggunakan composition
	type User struct {
		Name    string
		Age     int
		Address // embedded struct (composition)
	}

	// Membuat object User
	p := User{
		Name:  "Andi",
		Age:   25,
		Address: Address {
			City:	"Bandung",
			Province: "Jawa Barat",
		},
	}

	// Menampilkan data user
	fmt.Println("Nama  :", p.Name)
	fmt.Println("Umur  :", p.Age)
	fmt.Printf("Alamat : %v, %v\n", p.Address.City, p.Address.Province)
}