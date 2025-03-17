package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const UserSchemaName = "users"

type UserSchema struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	PubKey    string             `bson:"pubKey"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}
