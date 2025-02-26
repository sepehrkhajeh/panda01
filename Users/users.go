package Users

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	NIP05  string             `bson:"nip05"`
	PubKey string             `bson:"pubkey"`
	Relays []string           `bson:"relays"`
}

type UserJs struct {
	Name   string   `json:"name"`
	PubKey string   `json:"pubkey"`
	Relays []string `json:"relays"`
}

// This `User` struct represents a user in your application, with fields for the user's ID, name, NIP-05 identifier, public key, and relays. The `primitive.ObjectID` type is used for the ID field to ensure it is compatible with MongoDB's `_id` field. The `bson` tags are used to specify how the fields should be serialized to and from BSON, MongoDB's binary JSON format.
