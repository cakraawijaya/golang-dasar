package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score int = 95
	var text string = strconv.Itoa(score) // int to string

	fmt.Println("Nilai ujian:", text)
}