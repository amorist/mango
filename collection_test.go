package mango

import (
	"reflect"
	"testing"
)

func TestCollection_Find(t *testing.T) {
	type args struct {
		filter interface{}
	}
	tests := []struct {
		name string
		c    *Collection
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Find(tt.args.filter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}
