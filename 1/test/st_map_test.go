package test

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/n0fr1/-n0fr1-go-homework-2-reflex-7hw-/1/fillStructFromMap"
)

type DataSt struct {
	strExpect Person
	strResult Person
	mp        map[string]interface{}
}

type Person struct {
	Name string
	Age  int
}

func fillDataForTest() (D DataSt) {

	var AllData DataSt

	AllData.strExpect = struct {
		Name string
		Age  int
	}{
		Name: "Mari",
		Age:  26,
	}

	AllData.mp = map[string]interface{}{
		"Name":    "Maria",
		"Age":     26,
		"married": false,
	}

	return AllData
}

func TestStructToMap(t *testing.T) {

	AllData := fillDataForTest()

	err := fillStructFromMap.MapToStruct(&AllData.strResult, AllData.mp)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	if !(AllData.strResult.Name == AllData.strExpect.Name) { //comparing two structures
		t.Error("Name expected", AllData.strExpect.Name, "Got", AllData.strResult.Name)
	}

	if !(AllData.strResult.Age == AllData.strExpect.Age) {
		t.Error("Age expected", AllData.strExpect.Age, "Got", AllData.strResult.Age)
	}

}

func TestReflex(t *testing.T) {

	AllData := fillDataForTest()

	err := fillStructFromMap.MapToStruct(&AllData.strResult, AllData.mp)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	t.Run("test", func(t *testing.T) {
		if !reflect.DeepEqual(AllData.strExpect, AllData.strResult) {
			t.Error("Expected", AllData.strExpect, "got", AllData.strResult)
		}

	})

}
