package mango

import "go.mongodb.org/mongo-driver/mongo"

// Collection mongo-driver collection
type Collection struct {
	collection *mongo.Collection
}

// Find finds docs by given filter
func (c *Collection) Find(filter interface{}) *Session {
	return &Session{filter: filter, collection: c.collection}
}
