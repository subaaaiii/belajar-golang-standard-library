package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2002, time.April, 14, 10, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formater := "2 January 2006"
	value := "14 April 2002"
	// value := "14-04-2002"

	// RFC3339
	valueTime, err := time.Parse(formater, value)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())

	var duration1 time.Duration = time.Second * 100
	var duration2 time.Duration = time.Millisecond * 1000
	duration3 := duration1 - duration2

	fmt.Println(duration3)
	fmt.Printf("Duration: %d\n", duration3)

}
