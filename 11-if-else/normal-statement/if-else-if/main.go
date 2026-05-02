package main

import "fmt"

func main() {
	nilai := 75
	if nilai >= 90 {
		fmt.Printf("Nilai: %v --> Grade A", nilai)
	} else if nilai >= 75 {
		fmt.Printf("Nilai: %v --> Grade B", nilai)
	} else {
		fmt.Printf("Nilai: %v --> Grade C", nilai)
	}
}