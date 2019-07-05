package mango

import (
	"reflect"
	"testing"
)

func TestSession_SetPoolLimit(t *testing.T) {
	type args struct {
		limit uint16
	}
	var session *Session
	session = New("mongodb://127.0.0.1")
	tests := []struct {
		name string
		s    *Session
		args args
	}{
		{"SetPoolLimit", session, args{limit: 10}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.SetPoolLimit(tt.args.limit)
		})
	}
}

func TestSession_New(t *testing.T) {
	type args struct {
		uri string
	}
	var session *Session
	tests := []struct {
		name string
		s    *Session
		args args
		want *Session
	}{
		{"new session", session, args{uri: "mongodb://127.0.0.1"}, New("mongodb://127.0.0.1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Session.New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSession_Connect(t *testing.T) {
	var session *Session
	session = New("mongodb://127.0.0.1")
	tests := []struct {
		name    string
		s       *Session
		wantErr bool
	}{
		{"Connect", session, false},
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
		result []interface{}
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
