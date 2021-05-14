package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/n0fr1/-n0fr1-go-homework-2-reflex-7hw-/1/convertMapToStruct"
)

func TestStructToMap(t *testing.T) {

	st := struct {
		Name string
		Age  int
	}{
		Name: "Maria",
		Age:  25,
	}

	mp := map[string]interface{}{
		"Name":    "Maria",
		"Age":     25,
		"married": false,
	}

	err := convertMapToStruct.MapToStruct(&st, mp)
	if err != nil {
		fmt.Printf("error: %s", err)
		os.Exit(1)
	}

	if !(st.Name == mp["Name"]) {
		t.Error("st.Name expected", mp["Name"], "Got", st.Name)
	}
	if !(st.Age == mp["Age"]) {
		t.Error("st.Age expected", mp["Age"], "Got", st.Age)
	}

}
