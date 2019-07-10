package bson

import (
	"encoding/hex"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Zeroer allows custom struct types to implement a report of zero
// state. All struct types that don't implement Zeroer or where IsZero
// returns false are considered to be not zero.
type Zeroer interface {
	IsZero() bool
}

// D represents a BSON Document. This type can be used to represent BSON in a concise and readable
// manner. It should generally be used when serializing to BSON. For deserializing, the Raw or
// Document types should be used.
//
// Example usage:
//
// 		bson.D{{"foo", "bar"}, {"hello", "world"}, {"pi", 3.14159}}
//
// This type should be used in situations where order matters, such as MongoDB commands. If the
// order is not important, a map is more comfortable and concise.
type D = primitive.D

// E represents a BSON element for a D. It is usually used inside a D.
type E = primitive.E

// M is an unordered, concise representation of a BSON Document. It should generally be used to
// serialize BSON when the order of the elements of a BSON document do not matter. If the element
// order matters, use a D instead.
//
// Example usage:
//
// 		bson.M{"foo": "bar", "hello": "world", "pi": 3.14159}
//
// This type is handled in the encoders as a regular map[string]interface{}. The elements will be
// serialized in an undefined, random order, and the order will be different each time.
type M = primitive.M

// An A represents a BSON array. This type can be used to represent a BSON array in a concise and
// readable manner. It should generally be used when serializing to BSON. For deserializing, the
// RawArray or Array types should be used.
//
// Example usage:
//
// 		bson.A{"bar", "world", 3.14159, bson.D{{"qux", 12345}}}
//
type A = primitive.A

// ObjectID primitive.ObjectID
type ObjectID = primitive.ObjectID

// NewObjectID primitive.NewObjectID
var NewObjectID = primitive.NewObjectID

// ObjectIDHex primitive.ObjectIDFromHex
var ObjectIDHex = primitive.ObjectIDFromHex

// Marshal bson.Marshal
var Marshal = bson.Marshal

// IsObjectIDHex returns whether s is a valid hex representation of
// an ObjectId. See the ObjectIdHex function.
func IsObjectIDHex(s string) bool {
	if len(s) != 24 {
		return false
	}
	_, err := hex.DecodeString(s)
	return err == nil
}
