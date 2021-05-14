package main

import (
	"fmt"
	"reflect"
)

type str struct {
	field1 string
	field2 int
	field3 bool
}

func main() {
	example()
}
func change(in interface{}, values map[string]interface{}) interface{} {
	var err error
	_ = err
	switch in.(type) {
	case str:
		ps := reflect.ValueOf(in)
		for key, value := range values {
			f := ps.FieldByName(key)
			kind := reflect.ValueOf(value).Kind()
			fmt.Println(kind, reflect.ValueOf(value))
			if f.Kind() == kind {
				f.Set(value.(reflect.Value))
			}
		}
		return in
	default:
		return err
	}
}

func example() {
	a := str{
		field1: "stroka",
		field2: 126,
		field3: false,
	}
	values := map[string]interface{}{
		"field1": "strokaInter",
		"field2": 250,
		"field3": true,
	}
	fmt.Println(change(a, values))
}
