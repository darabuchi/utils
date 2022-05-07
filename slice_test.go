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
