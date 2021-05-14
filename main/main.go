package main

import (
	"errors"
	"fmt"
	"reflect"
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

	err := mapToStruct(&st, mp)

	fmt.Printf("STRUCT: %#v\nERR: %s\n", st, err)
}

func mapToStruct(in interface{}, values map[string]interface{}) error {

	valueOf := reflect.ValueOf(in)
	if valueOf.Kind() != reflect.Ptr {
		return errors.New("not a pointer to struct")
	}

	valueOf = valueOf.Elem()
	if valueOf.Kind() != reflect.Struct {
		return errors.New("not a pointer to struct")
	}

	valueType := valueOf.Type()
	for i := 0; i < valueType.NumField(); i++ {

		field := valueType.Field(i) // reflect.StructField
		fv := valueOf.Field(i)      // reflect.Value

		if val, ok := values[field.Name]; ok {

			mfv := reflect.ValueOf(val)
			if mfv.Kind() != fv.Kind() {
				return errors.New("incomatible type for " + field.Name)
			}

			fv.Set(mfv)
		}
	}

	return nil
}
