package mango

import (
	"context"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Session mongo session
type Session struct {
	client      *mongo.Client
	collection  *mongo.Collection
	maxPoolSize uint16
	db          string
	uri         string
	m           sync.RWMutex
	filter      interface{}
	limit       int64
	project     interface{}
	skip        int64
	sort        interface{}
}

// New session
func New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

// SetDB set db
func (s *Session) SetDB(db string) {
	s.m.Lock()
	s.db = db
	s.m.Unlock()
}

// C returns collection
func (s *Session) C(collection string) *Collection {
	if len(s.db) == 0 {
		s.db = "test"
	}
	d := &Database{database: s.client.Database(s.db)}
	return &Collection{collection: d.database.Collection(collection)}
}

// Collection returns collection
func (s *Session) Collection(collection string) *Collection {
	if len(s.db) == 0 {
		s.db = "test"
	}
	d := &Database{database: s.client.Database(s.db)}
	return &Collection{collection: d.database.Collection(collection)}
}

// SetPoolLimit set maxPoolSize
func (s *Session) SetPoolLimit(limit uint16) {
	s.m.Lock()
	s.maxPoolSize = limit
	s.m.Unlock()
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

// Limit limit
func (s *Session) Limit(limit int64) *Session {
	s.limit = limit
	return s
}

// Skip Skip
func (s *Session) Skip(skip int64) *Session {
	s.skip = skip
	return s
}

// Sort sort
func (s *Session) Sort(sort interface{}) *Session {
	s.sort = sort
	return s
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

// All find all
func (s *Session) All(result interface{}) error {
	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr {
		panic("result argument must be a slice address")
	}
	slicev := resultv.Elem()

	if slicev.Kind() == reflect.Interface {
		slicev = slicev.Elem()
	}
	if slicev.Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}

	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	cur, err := s.collection.Find(ctx, s.filter)
	defer cur.Close(ctx)
	if err != nil {
		return err
	}
	if err = cur.Err(); err != nil {
		return err
	}
	i := 0
	for cur.Next(ctx) {
		elemp := reflect.New(elemt)
		if err = bson.Unmarshal(cur.Current, elemp.Interface()); err != nil {
			return err
		}
		slicev = reflect.Append(slicev, elemp.Elem())
		i++
	}
	resultv.Elem().Set(slicev.Slice(0, i))
	return nil
}
