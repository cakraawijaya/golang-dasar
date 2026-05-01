package main

import "fmt"

func main() {
	hari := "Sabtu"
	fmt.Print("Input = ", hari)

	switch hari {
		case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
			fmt.Println(" --> Hari Kerja")
		case "Sabtu", "Minggu":
			fmt.Println(" --> Hari Libur")
		default:
			fmt.Println(" --> Hari tidak valid")
	}
}