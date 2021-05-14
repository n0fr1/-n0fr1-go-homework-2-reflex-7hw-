package convertMapToStruct

import (
	"errors"
	"reflect"
)

func MapToStruct(in interface{}, values map[string]interface{}) error {

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
