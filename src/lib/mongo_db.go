package lib

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoDB implements DBInterface for MongoDB
type MongoDB struct {
	Db *mongo.Database
}

func (db *MongoDB) Create(ctx context.Context, collection string, data interface{}) error {
	_, err := db.Db.Collection(collection).InsertOne(ctx, data)
	return err
}

func (db *MongoDB) Find(ctx context.Context, collection string, query interface{}, result interface{}) error {
	return db.Db.Collection(collection).FindOne(ctx, query).Decode(result)
}

func (db *MongoDB) Update(ctx context.Context, collection string, query interface{}, updateData interface{}) error {
	_, err := db.Db.Collection(collection).UpdateOne(ctx, query, updateData)
	return err
}

func (db *MongoDB) Delete(ctx context.Context, collection string, query interface{}) error {
	_, err := db.Db.Collection(collection).DeleteOne(ctx, query)
	return err
}
