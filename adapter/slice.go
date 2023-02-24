package adapter

import (
	"fmt"
	"github.com/pkg6/go-paginator"
	"reflect"
)

func NewSliceAdapter(source any) paginator.IAdapter {
	if isPtr(source) || !isSlice(source) {
		panic(fmt.Sprintf("expected slice but got %s", reflect.TypeOf(source).Kind()))
	}
	return &SliceAdapter{src: source}
}

type SliceAdapter struct {
	src any
}

func (s SliceAdapter) Length() (int64, error) {
	return int64(reflect.ValueOf(s.src).Len()), nil
}

func (s SliceAdapter) Slice(offset, length int64, dest any) error {
	va := reflect.ValueOf(s.src)
	fullSize := int64(va.Len())
	needSize := length + offset
	if fullSize < needSize {
		length = fullSize - offset
	}
	lengthInt := int(length)
	if lengthInt <= 0 {
		lengthInt = 0
	}
	if err := makeSlice(dest, lengthInt, lengthInt); err != nil {
		return err
	}
	// 超出切片可以切割的范围
	if lengthInt == 0 {
		return nil
	}
	//防止切片需要切割的过多，导致的失败
	if needSize > fullSize {
		needSize = fullSize
	}
	vs := va.Slice(int(offset), int(needSize))
	vt := reflect.ValueOf(dest).Elem()
	for i := 0; i < vs.Len(); i++ {
		vt.Index(i).Set(reflect.ValueOf(vs.Index(i).Interface()))
	}
	return nil
}
