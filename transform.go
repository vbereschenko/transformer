package transformer

import (
	"reflect"
	"errors"
)

// Transform moves field from one structure to another
// second structure should be provided as pointer
//
// source - structure pointer or structure which we take fields from
// target - structure we move fields to
func Transform(source, target interface{}) error {
	typeDescriptor := reflect.TypeOf(source)
	valueDescriptor := reflect.ValueOf(source)
	if valueDescriptor.Kind() == reflect.Ptr {
		valueDescriptor = valueDescriptor.Elem()
		typeDescriptor = typeDescriptor.Elem()
	}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target should be pointer")
	}
	targetValue = targetValue.Elem()
	for i := 0; i< typeDescriptor.NumField(); i++ {
		fieldType := typeDescriptor.Field(i)
		fieldValue := valueDescriptor.Field(i)

		targetField := targetValue.FieldByName(fieldType.Name)
		if targetField.Kind() == fieldValue.Kind() {
			targetField.Set(fieldValue)
		}
	}

	return nil
}
