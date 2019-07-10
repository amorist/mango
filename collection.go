package mango

import (
	"context"
	"time"

	"github.com/amorist/mango/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection mongo-driver collection
type Collection struct {
	collection *mongo.Collection
}

// Find finds docs by given filter
func (c *Collection) Find(filter interface{}) *Session {
	return &Session{filter: filter, collection: c.collection}
}

// Insert inserts a single document into the collection.
func (c *Collection) Insert(document interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	if _, err = c.collection.InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}

// InsertAll inserts the provided documents.
func (c *Collection) InsertAll(documents []interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	if _, err = c.collection.InsertMany(ctx, documents); err != nil {
		return err
	}
	return nil
}

// Update updates a single document in the collection.
func (c *Collection) Update(selector interface{}, update interface{}, upsert ...bool) error {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error

	opt := options.Update()
	for _, arg := range upsert {
		if arg {
			opt.SetUpsert(arg)
		}
	}

	if _, err = c.collection.UpdateOne(ctx, selector, update, opt); err != nil {
		return err
	}
	return nil
}

// UpdateID updates a single document in the collection by id
func (c *Collection) UpdateID(id interface{}, update interface{}) error {
	return c.Update(bson.M{"_id": id}, update)
}

// UpdateAll updates multiple documents in the collection.
func (c *Collection) UpdateAll(selector interface{}, update interface{}, upsert ...bool) (*mongo.UpdateResult, error) {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error

	opt := options.Update()
	for _, arg := range upsert {
		if arg {
			opt.SetUpsert(arg)
		}
	}

	var updateResult *mongo.UpdateResult
	if updateResult, err = c.collection.UpdateMany(ctx, selector, update, opt); err != nil {
		return updateResult, err
	}
	return updateResult, nil
}

// Remove deletes a single document from the collection.
func (c *Collection) Remove(selector interface{}) error {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	if _, err = c.collection.DeleteOne(ctx, selector); err != nil {
		return err
	}
	return nil
}

// RemoveID deletes a single document from the collection by id.
func (c *Collection) RemoveID(id interface{}) error {
	return c.Remove(id)
}

// RemoveAll deletes multiple documents from the collection.
func (c *Collection) RemoveAll(selector interface{}) error {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error

	if _, err = c.collection.DeleteMany(ctx, selector); err != nil {
		return err
	}
	return nil
}

// Count gets the number of documents matching the filter.
func (c *Collection) Count(selector interface{}) (int64, error) {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var count int64
	count, err = c.collection.CountDocuments(ctx, selector)
	return count, err
}
