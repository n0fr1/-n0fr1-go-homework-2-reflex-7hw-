package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/n0fr1/-n0fr1-go-homework-2-reflex-7hw-/1/fillStructFromMap"
)

type Person struct {
	Name     string
	LastName string
	Age      int
}

func TestStructToMap(t *testing.T) {

	var strResult Person

	strExpect := struct {
		Name string
		Age  int
	}{
		Name: "Maria",
		Age:  26,
	}

	mp := map[string]interface{}{
		"Name":    "Maria",
		"Age":     26,
		"married": false,
	}

	err := fillStructFromMap.MapToStruct(&strResult, mp)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	if !(strResult.Name == strExpect.Name) { //comparing two structures
		t.Error("Name expected", strExpect.Name, "Got", strResult.Name)
	}

	if !(strResult.Age == strExpect.Age) {
		t.Error("Age expected", strExpect.Age, "Got", strResult.Age)
	}

}
