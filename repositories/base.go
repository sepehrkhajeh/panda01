package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseRepository struct {
	Client       *mongo.Client
	DBName       string
	Collection   string
	QueryTimeout time.Duration
}

func NewBaseRepository(client *mongo.Client, dbName, collection string, queryTimeout time.Duration) *BaseRepository {
	return &BaseRepository{
		Client:       client,
		DBName:       dbName,
		Collection:   collection,
		QueryTimeout: queryTimeout,
	}
}

func (r *BaseRepository) InsertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)
	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()
	return collection.InsertOne(ctx, document)

}

// func (r *BaseRepository) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
// 	collection := r.Client.Database(r.DBName).Collection(r.Collection)
// 	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
// 	defer cancel()

// 	err := collection.FindOne(ctx, filter).Decode(result)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
func (r *BaseRepository) FindOne(ctx context.Context, filter interface{}, result interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	if r.QueryTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, r.QueryTimeout)
		defer cancel()
	}
	return collection.FindOne(ctx, filter).Decode(result)
}

func (r *BaseRepository) FindAll(ctx context.Context, filter interface{}, results interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	// بررسی مقدار nil بودن cursor
	if cursor == nil {
		return mongo.ErrNoDocuments
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, results)
}

// func (r *BaseRepository) insertOne(ctx context.Context, document interface{}) (*mongo.InsertOneResult, error) {
// 	collection := r.Client.Database(r.DBName).Collection(r.Collection)
// 	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
// 	defer cancel()
// 	return collection.InsertOne(ctx, document)
// }

func (r *BaseRepository) FindByField(ctx context.Context, field string, value interface{}, result interface{}) error {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)
	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()
	filter := bson.M{field: value}
	err := collection.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseRepository) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)

	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return result, err
	}

	if result.MatchedCount == 0 {
		return result, mongo.ErrNoDocuments
	}

	return result, nil
}

func (r *BaseRepository) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)
	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return result.DeletedCount, nil
}

func (r *BaseRepository) CountDocuments(ctx context.Context, filter interface{}) (int64, error) {
	collection := r.Client.Database(r.DBName).Collection(r.Collection)
	ctx, cancel := context.WithTimeout(ctx, r.QueryTimeout)
	defer cancel()
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
