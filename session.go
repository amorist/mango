package mango

import (
	"context"
	"reflect"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Session mongo session
type Session struct {
	client      *mongo.Client
	collection  *mongo.Collection
	maxPoolSize uint64
	db          string
	uri         string
	m           sync.RWMutex
	filter      interface{}
	limit       *int64
	project     interface{}
	skip        *int64
	sort        interface{}
}

// New session
//
// Relevant documentation:
//
// 		https://docs.mongodb.com/manual/reference/connection-string/
func New(uri string) *Session {
	session := &Session{
		uri: uri,
	}
	return session
}

// C Collection alias
func (s *Session) C(collection string) *Collection {
	s.m.Lock()
	defer s.m.Unlock()
	if len(s.db) == 0 {
		s.db = "test"
	}
	d := &Database{database: s.client.Database(s.db)}
	return &Collection{collection: d.database.Collection(collection)}
}

// Collection returns collection
func (s *Session) Collection(collection string) *Collection {
	s.m.Lock()
	defer s.m.Unlock()
	if len(s.db) == 0 {
		s.db = "test"
	}
	d := &Database{database: s.client.Database(s.db)}
	return &Collection{collection: d.database.Collection(collection)}
}

// SetPoolLimit specifies the max size of a server's connection pool.
func (s *Session) SetPoolLimit(limit uint64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.maxPoolSize = limit
}

// Connect mongo client
func (s *Session) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI(s.uri)
	opt.SetMaxPoolSize(s.maxPoolSize)

	client, err := mongo.NewClient(opt)
	if err != nil {
		return err
	}

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	s.client = client
	return nil
}

// Ping verifies that the client can connect to the topology.
// If readPreference is nil then will use the client's default read
// preference.
func (s *Session) Ping() error {
	return s.client.Ping(context.TODO(), readpref.Primary())
}

// Client return mongo Client
func (s *Session) Client() *mongo.Client {
	return s.client
}

// DB returns a value representing the named database.
func (s *Session) DB(db string) *Database {
	s.m.Lock()
	defer s.m.Unlock()
	return &Database{database: s.client.Database(db)}
}

// Limit specifies a limit on the number of results.
// A negative limit implies that only 1 batch should be returned.
func (s *Session) Limit(limit int64) *Session {
	s.limit = &limit
	return s
}

// Skip specifies the number of documents to skip before returning.
// For server versions < 3.2, this defaults to 0.
func (s *Session) Skip(skip int64) *Session {
	s.skip = &skip
	return s
}

// Sort specifies the order in which to return documents.
func (s *Session) Sort(sort interface{}) *Session {
	s.sort = sort
	return s
}

// One returns up to one document that matches the model.
func (s *Session) One(result interface{}) error {
	var err error
	data, err := s.collection.FindOne(context.TODO(), s.filter).DecodeBytes()

	if err != nil {
		return err
	}

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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	var err error

	opt := options.Find()

	if s.sort != nil {
		opt.SetSort(s.sort)
	}

	if s.limit != nil {
		opt.SetLimit(*s.limit)
	}

	if s.skip != nil {
		opt.SetSkip(*s.skip)
	}

	cur, err := s.collection.Find(ctx, s.filter, opt)
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
