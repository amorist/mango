package mango

import (
	"context"
	"mango/bson"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

// Collection mongo-driver collection
type Collection struct {
	collection *mongo.Collection
}

// Find finds docs by given filter
func (c *Collection) Find(filter interface{}) *Session {
	return &Session{filter: filter, collection: c.collection}
}

// Insert one
func (c *Collection) Insert(document interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	if _, err = c.collection.InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}

// InsertAll all
func (c *Collection) InsertAll(documents []interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	if _, err = c.collection.InsertMany(ctx, documents); err != nil {
		return err
	}
	return nil
}

// Update update one
func (c *Collection) Update(selector interface{}, update interface{}) error {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error

	if _, err = c.collection.UpdateOne(ctx, selector, update); err != nil {
		return err
	}
	return nil
}

// UpdateID update by id
func (c *Collection) UpdateID(id interface{}, update interface{}) error {
	return c.Update(bson.M{"_id": id}, update)
}

// UpdateAll update all
func (c *Collection) UpdateAll(selector interface{}, update interface{}) (*mongo.UpdateResult, error) {
	if selector == nil {
		selector = bson.D{}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var err error
	var updateResult *mongo.UpdateResult
	if updateResult, err = c.collection.UpdateMany(ctx, selector, update); err != nil {
		return updateResult, err
	}
	return updateResult, nil
}

// Remove by selector
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

// RemoveID remove by id
func (c *Collection) RemoveID(id interface{}) error {
	return c.Remove(id)
}

// RemoveAll remove all
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

// Count count
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
