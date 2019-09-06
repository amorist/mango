package mango

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Database mongo-driver database
type Database struct {
	database *mongo.Database
}

// CollectionNames returns the collection names present in database.
func (d *Database) CollectionNames() (names []string, err error) {
	names, err = d.database.ListCollectionNames(context.TODO(), options.ListCollectionsOptions{})
	return
}

// C returns collection.
func (d *Database) C(collection string) *Collection {
	return &Collection{collection: d.database.Collection(collection)}
}

// Collection returns collection.
func (d *Database) Collection(collection string) *Collection {
	return &Collection{collection: d.database.Collection(collection)}
}
