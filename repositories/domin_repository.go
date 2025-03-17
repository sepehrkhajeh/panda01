package repositories

import (
	"Panda/schemas"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DomainRepository struct {
	*BaseRepository
}

func NewDomainRepository(client *mongo.Client, dbname string, timeout time.Duration) *DomainRepository {
	return &DomainRepository{
		BaseRepository: NewBaseRepository(client, dbname, schemas.DomainSchemaName, timeout),
	}
}

func (r *DomainRepository) Add(ctx context.Context, d *schemas.DomainSchema) (*mongo.InsertOneResult, error) {
	d.ID = primitive.NewObjectID()
	d.CreatedAt = time.Now()
	d.UpdatedAt = time.Now()
	return r.InsertOne(ctx, d)

}

func (r *DomainRepository) Delete(ctx context.Context, filter interface{}) (int64, error) {
	result, err := r.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (r *DomainRepository) Update(ctx context.Context, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	return r.UpdateOne(ctx, filter, update)
}

func (r *DomainRepository) GetByFeild(ctx context.Context, fieldName string, value interface{}) (*schemas.DomainSchema, error) {
	domain := new(schemas.DomainSchema)
	err := r.FindOne(ctx, bson.M{fieldName: value}, domain)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return domain, nil

}

func (r *DomainRepository) GetByID(ctx context.Context, id string) (*schemas.DomainSchema, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	domain := new(schemas.DomainSchema)

	err = r.FindOne(ctx, bson.M{"_id": objectID}, domain)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return domain, nil
}

func (r *DomainRepository) GetAll(ctx context.Context, filter interface{}) (*[]schemas.DomainSchema, error) {
	results := new([]schemas.DomainSchema)
	err := r.FindAll(ctx, filter, results)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return &[]schemas.DomainSchema{}, nil
		}

		return nil, err
	}

	return results, nil
}
