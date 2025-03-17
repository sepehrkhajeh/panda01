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

type IdentifierRepository struct {
	*BaseRepository
}

func NewIdentifierRepository(client *mongo.Client, dbName string, timeout time.Duration) *IdentifierRepository {
	return &IdentifierRepository{
		BaseRepository: NewBaseRepository(client, dbName, schemas.IdentifierSchemaName, timeout),
	}
}

func (r *IdentifierRepository) Add(ctx context.Context, d schemas.IdentifierSchema) (*mongo.InsertOneResult, error) {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()

	return r.InsertOne(ctx, d)
}

func (r *IdentifierRepository) GetAll(ctx context.Context, filter interface{}) (*[]schemas.IdentifierSchema, error) {
	results := new([]schemas.IdentifierSchema)
	err := r.FindAll(ctx, filter, results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &[]schemas.IdentifierSchema{}, nil
		}
		return nil, err
	}
	return results, nil
}

func (r *IdentifierRepository) Update(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.UpdateOne(ctx, filter, update)
}

func (r *IdentifierRepository) Delete(ctx context.Context, filter interface{}) (int64, error) {
	result, err := r.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *IdentifierRepository) GetByFeild(ctx context.Context, fieldName string, value interface{}) (*schemas.IdentifierSchema, error) {
	result := new(schemas.IdentifierSchema)
	err := r.FindOne(ctx, bson.M{fieldName: value}, result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return result, nil
}
