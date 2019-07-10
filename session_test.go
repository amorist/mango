package mango

import (
	"reflect"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
)

func TestNew(t *testing.T) {
	type args struct {
		uri string
	}
	session := New("mongodb://127.0.0.1")

	tests := []struct {
		name string
		args args
		want *Session
	}{
		{"new", args{uri: "mongodb://127.0.0.1"}, session},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_SetDB(t *testing.T) {
	type args struct {
		db string
	}
	session := New("mongodb://127.0.0.1")
	tests := []struct {
		name string
		s    *Session
		args args
	}{
		{"set db", session, args{db: "test"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetDB(tt.args.db)
		})
	}
}

func TestSession_C(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.C(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.C() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Collection(t *testing.T) {
	type args struct {
		collection string
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Collection
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Collection(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Collection() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_SetPoolLimit(t *testing.T) {
	type args struct {
		limit uint16
	}
	tests := []struct {
		name string
		s    *Session
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetPoolLimit(tt.args.limit)
		})
	}
}

func TestSession_Connect(t *testing.T) {
	tests := []struct {
		name    string
		s       *Session
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Connect(); (err != nil) != tt.wantErr {
				t.Errorf("Session.Connect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_Ping(t *testing.T) {
	var session *Session
	session = New("mongodb://127.0.0.1")
	session.Connect()
	tests := []struct {
		name    string
		s       *Session
		wantErr bool
	}{
		{"ping", session, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("Session.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_Client(t *testing.T) {
	tests := []struct {
		name string
		s    *Session
		want *mongo.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Client(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Client() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_DB(t *testing.T) {
	type args struct {
		db string
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Database
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.DB(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.DB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Limit(t *testing.T) {
	type args struct {
		limit int64
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Limit(tt.args.limit); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Limit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Skip(t *testing.T) {
	type args struct {
		skip int64
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Skip(tt.args.skip); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Skip() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Sort(t *testing.T) {
	type args struct {
		sort interface{}
	}
	tests := []struct {
		name string
		s    *Session
		args args
		want *Session
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.Sort(tt.args.sort); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_One(t *testing.T) {
	type args struct {
		result interface{}
	}
	tests := []struct {
		name    string
		s       *Session
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.One(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Session.One() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSession_All(t *testing.T) {
	type args struct {
		result interface{}
	}
	tests := []struct {
		name    string
		s       *Session
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.All(tt.args.result); (err != nil) != tt.wantErr {
				t.Errorf("Session.All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
