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
	sourceTypeDescriptor := reflect.TypeOf(source)
	sourceValueDescriptor := reflect.ValueOf(source)
	if sourceValueDescriptor.Kind() == reflect.Ptr {
		sourceValueDescriptor = sourceValueDescriptor.Elem()
		sourceTypeDescriptor = sourceTypeDescriptor.Elem()
	}

	targetValueDescriptor := reflect.ValueOf(target)
	if targetValueDescriptor.Kind() != reflect.Ptr {
		return errors.New("target should be pointer")
	}
	targetDescriptor := reflect.TypeOf(target).Elem()
	targetValueDescriptor = targetValueDescriptor.Elem()
	var name string
	for i := 0; i< targetDescriptor.NumField(); i++ {
		fieldValue := targetDescriptor.Field(i)
		if value, found := fieldValue.Tag.Lookup("fromField"); found {
			name = value
		} else {
			name = targetDescriptor.Field(i).Name
		}
		sourceField := sourceValueDescriptor.FieldByName(name)
		if sourceField.Kind() == targetValueDescriptor.Field(i).Kind() {
			targetValueDescriptor.Field(i).Set(sourceField.Convert(targetValueDescriptor.Field(i).Type()))
		}
	}

	return nil
}
