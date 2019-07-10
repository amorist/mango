package mango

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
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

func TestCollection_Insert(t *testing.T) {
	type args struct {
		document interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Insert(tt.args.document); (err != nil) != tt.wantErr {
				t.Errorf("Collection.Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_InsertAll(t *testing.T) {
	type args struct {
		documents []interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.InsertAll(tt.args.documents); (err != nil) != tt.wantErr {
				t.Errorf("Collection.InsertAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_Update(t *testing.T) {
	type args struct {
		selector interface{}
		update   interface{}
		upsert   []bool
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Update(tt.args.selector, tt.args.update, tt.args.upsert...); (err != nil) != tt.wantErr {
				t.Errorf("Collection.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_UpdateID(t *testing.T) {
	type args struct {
		id     interface{}
		update interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UpdateID(tt.args.id, tt.args.update); (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_UpdateAll(t *testing.T) {
	type args struct {
		selector interface{}
		update   interface{}
		upsert   []bool
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		want    *mongo.UpdateResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.UpdateAll(tt.args.selector, tt.args.update, tt.args.upsert...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.UpdateAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Collection.UpdateAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCollection_Remove(t *testing.T) {
	type args struct {
		selector interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Remove(tt.args.selector); (err != nil) != tt.wantErr {
				t.Errorf("Collection.Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_RemoveID(t *testing.T) {
	type args struct {
		id interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.RemoveID(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Collection.RemoveID() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_RemoveAll(t *testing.T) {
	type args struct {
		selector interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.RemoveAll(tt.args.selector); (err != nil) != tt.wantErr {
				t.Errorf("Collection.RemoveAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCollection_Count(t *testing.T) {
	type args struct {
		selector interface{}
	}
	tests := []struct {
		name    string
		c       *Collection
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Count(tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("Collection.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Collection.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}
