package mango

import (
	"context"
	"sync"
	"time"

	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Session mongo session
type Session struct {
	client      *mongo.Client
	collection  *mongo.Collection
	maxPoolSize uint16
	uri         string
	m           sync.RWMutex
	filter      interface{}
	limit       int64
	project     interface{}
	skip        int64
	sort        interface{}
}

// SetPoolLimit set maxPoolSize
func (s *Session) SetPoolLimit(limit uint16) {
	s.m.Lock()
	s.maxPoolSize = limit
	s.m.Unlock()
}

// New session
func (s *Session) New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

// Connect mongo
func (s *Session) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI(s.uri)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		return err
	}
	s.client = client
	return nil
}

// Ping -
func (s *Session) Ping() error {
	return s.Ping()
}

// DB db
func (s *Session) DB(db string) *Database {
	return &Database{database: s.client.Database(db)}
}

// One find one
func (s *Session) One(result interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	data, err := s.collection.FindOne(ctx, s.filter).DecodeBytes()
	err = bson.Unmarshal(data, result)
	return err
}
