//1. Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
//values map[string]interface{} (key - название поля структуры, которому нужно присвоить value
//этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
//пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции
//тесты (чем больше, тем лучше - зачтется в плюс).

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
