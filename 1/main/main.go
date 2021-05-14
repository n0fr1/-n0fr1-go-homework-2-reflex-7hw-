package main

import (
	"fmt"

	"github.com/n0fr1/-n0fr1-go-homework-2-reflex-7hw-/1/fillStructFromMap"
)

type Person struct {
	Name     string
	LastName string
	Age      int
}

func main() {

	var stPerson Person

	mp := map[string]interface{}{
		"Name":    "John",
		"Age":     30,
		"married": true,
	}

	err := fillStructFromMap.MapToStruct(&stPerson, mp)

	fmt.Printf("STRUCT: %#v\nERR: %s\n", stPerson, err)
}
