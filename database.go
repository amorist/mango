package mango

import "go.mongodb.org/mongo-driver/mongo"

// Database mongo-driver database
type Database struct {
	database *mongo.Database
}

// C returns collection
func (d *Database) C(collection string) *Collection {
	return &Collection{collection: d.database.Collection(collection)}
}
