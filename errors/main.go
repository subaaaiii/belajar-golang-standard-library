package main

import (
	"errors"
	"fmt"
)

var (
	validationError = errors.New("validation error")
	notFoundError   = errors.New("Not found error")
)

func getByID(id string) error {
	if id == "" {
		return validationError
	}
	if id != "subai" {
		return notFoundError
	}

	return nil
}

func main() {
	err := getByID("")

	if err != nil {
		if errors.Is(err, validationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, notFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error")
		}
	} else {
		fmt.Println("sukses")
	}

}
