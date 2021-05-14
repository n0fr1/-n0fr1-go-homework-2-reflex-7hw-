package main

import (
	"fmt"

	"github.com/n0fr1/-n0fr1-go-homework-2-reflex-7hw-/1/convertMapToStruct"
)

type St struct {
	Name     string
	LastName string
	Age      int
}

func main() {

	var st St

	mp := map[string]interface{}{
		"Name":    "John",
		"Age":     30,
		"married": true,
	}

	err := convertMapToStruct.MapToStruct(&st, mp)

	fmt.Printf("STRUCT: %#v\nERR: %s\n", st, err)
}
