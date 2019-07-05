package mango

import (
	"reflect"
	"testing"
)

func TestDatabase_C(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name string
		d    *Database
		args args
		want *Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.d.C(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Database.C() = %v, want %v", got, tt.want)
			}
		})
	}
}
