package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/sepehrkhajeh/panda01/schemas"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepasitory struct {
	*BaseRepository
}

func NewUserRepository(client *mongo.Client, dbName string, timeout time.Duration) *UserRepasitory {
	return &UserRepasitory{
		BaseRepository: NewBaseRepository(client, dbName, schemas.UserSchemaName, timeout),
	}
}

func (r *UserRepasitory) GetByFeild(ctx context.Context, field string, value interface{}) (*schemas.UserSchema, error) {
	result := new(schemas.UserSchema)
	err := r.FindOne(ctx, bson.M{field: value}, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}

func (r *UserRepasitory) GetAll(ctx context.Context, filter interface{}) (*[]schemas.UserSchema, error) {
	result := new([]schemas.UserSchema)
	err := r.FindAll(ctx, filter, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *UserRepasitory) Delete(ctx context.Context, filter interface{}) (int64, error) {
	result, err := r.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *UserRepasitory) Update(ctx context.Context, item, update interface{}) (*mongo.UpdateResult, error) {
	updatedItem, err := r.UpdateOne(ctx, item, update)
	if err != nil {
		return updatedItem, err
	}
	return updatedItem, nil
}

func (r *UserRepasitory) Add(ctx context.Context, d *schemas.UserSchema) (*mongo.InsertOneResult, error) {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return r.InsertOne(ctx, d)
}
