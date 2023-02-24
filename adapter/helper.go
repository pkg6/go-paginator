package adapter

import (
	"fmt"
	"reflect"
)

func isPtr(data any) bool {
	t := reflect.TypeOf(data)
	return t.Kind() == reflect.Ptr
}
func isSlice(data any) bool {
	v := reflect.Indirect(reflect.ValueOf(data))
	return v.Kind() == reflect.Slice
}
func makeSlice(data interface{}, length, cap int) error {
	if !isPtr(data) {
		return fmt.Errorf("expected to be a ptr but got %T", data)
	}
	if !isSlice(data) {
		return fmt.Errorf("expected to be a slice pointer but got %T", data)
	}
	ind := reflect.Indirect(reflect.ValueOf(data))
	typ := reflect.TypeOf(ind.Interface())
	reflect.ValueOf(data).Elem().Set(reflect.MakeSlice(typ, length, cap))
	return nil
}
