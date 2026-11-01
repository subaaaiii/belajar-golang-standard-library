package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Subairi ahmad", "Subai"))
	fmt.Println(strings.Split("Subairi ahmad", " "))
	fmt.Println(strings.ToLower("Subairi ahmad"))
	fmt.Println(strings.ToUpper("Subairi ahmad"))
	fmt.Println(strings.Trim("      Subairi ahmad     ", " "))
	fmt.Println(strings.ReplaceAll("Subairi ahmad Subairi", "Subairi", "eko"))
}
