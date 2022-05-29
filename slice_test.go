package utils

import (
	"reflect"
	"testing"
)

func TestPluckUint64(t *testing.T) {
	type item1 struct {
		Id uint64
	}
	
	type item2 struct {
		Id uint32
	}
	
	type args struct {
		list      interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "",
			args: args{
				list: []item1{
					{
						1,
					},
					{
						2,
					},
				},
				fieldName: "Id",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PluckUint64(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PluckUint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeyByV2(t *testing.T) {
	type item struct {
		Id   int
		Name string
	}
	
	type args struct {
		list      interface{}
		fieldName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "",
			args: args{
				list: []*item{
					{
						Id:   1,
						Name: "a",
					},
					{
						Id:   2,
						Name: "b",
					},
					{
						Id:   3,
						Name: "c",
					},
				},
				fieldName: "Id",
			},
			want: map[int]*item{
				1: {
					Id:   1,
					Name: "a",
				},
				2: {
					Id:   2,
					Name: "b",
				},
				3: {
					Id:   3,
					Name: "c",
				},
			},
		},
		{
			name: "",
			args: args{
				list:      []*item{},
				fieldName: "Id",
			},
			want: map[int]*item{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeyBy(tt.args.list, tt.args.fieldName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeyBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
