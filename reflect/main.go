package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Age     int    `required:"true" max:"10"`
}

func ReadField(value any) {
	var valueType reflect.Type = reflect.TypeOf(value)
	fmt.Println("type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println("required:", valueField.Tag.Get("required"))
		fmt.Println("max:", valueField.Tag.Get("max"))
	}

}

func IsValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result

}

func main() {
	ReadField(Sample{"Eko"})
	ReadField(Person{"Eko", "", 12})

	fmt.Println(IsValid(Person{"subairi", "jakarta", 12}))
}
