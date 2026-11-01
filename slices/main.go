package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"john", "paul", "george", "ringo"}
	values := []int{10, 12, 9, 7, 13}
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(names, "paul"))
	fmt.Println(slices.Contains(names, "eko"))
	fmt.Println(slices.Index(names, "eko"))
	fmt.Println(slices.Index(names, "eko"))
}
