package adapter

import (
	"github.com/pkg6/go-paginator"
	"reflect"
	"testing"
)

func TestNewSliceAdapter(t *testing.T) {
	type args struct {
		source any
	}
	var tests []struct {
		name string
		args args
		want paginator.IAdapter
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSliceAdapter(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSliceAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceAdapter_Length(t *testing.T) {
	type fields struct {
		src any
	}
	var tests []struct {
		name    string
		fields  fields
		want    int64
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SliceAdapter{
				src: tt.fields.src,
			}
			got, err := s.Length()
			if (err != nil) != tt.wantErr {
				t.Errorf("Length() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Length() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSliceAdapter_Slice(t *testing.T) {
	type fields struct {
		src any
	}
	type args struct {
		offset int64
		length int64
		dest   any
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := SliceAdapter{
				src: tt.fields.src,
			}
			if err := s.Slice(tt.args.offset, tt.args.length, tt.args.dest); (err != nil) != tt.wantErr {
				t.Errorf("Slice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
