package lib

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB implements DBInterface for MongoDB
type MongoDB struct {
	Db *mongo.Database
}

type MongoQueryBuilder interface {
	Collection(name string) MongoQueryBuilder
	Filter(filter bson.M) MongoQueryBuilder
	Projection(fields bson.M) MongoQueryBuilder
	Build() (collection string, filter bson.M, projection bson.M)
}

type mongoQueryBuilder struct {
	collection string
	projection bson.M
	filter     bson.M
	joins      []bson.M // used for $lookup in aggregation
	args       []interface{}
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

func NewMongoQueryBuilder() QueryBuilder {
	return &mongoQueryBuilder{
		projection: bson.M{},
		filter:     bson.M{},
		joins:      []bson.M{},
	}
}

func (qb *mongoQueryBuilder) From(collection string) QueryBuilder {
	qb.collection = collection
	return qb
}

func (qb *mongoQueryBuilder) Select(fields ...string) QueryBuilder {
	for _, field := range fields {
		qb.projection[field] = 1
	}
	return qb
}

func (qb *mongoQueryBuilder) Join(join string, condition string) QueryBuilder {
	// assume simple $lookup style join with "foreignCollection.localField=foreignField"
	// e.g., "orders.user_id=users._id"
	var localField, foreignField string
	fmt.Sscanf(condition, "%s=%s", &localField, &foreignField)

	qb.joins = append(qb.joins, bson.M{
		"$lookup": bson.M{
			"from":         join,
			"localField":   localField,
			"foreignField": foreignField,
			"as":           join,
		},
	})
	return qb
}

func (qb *mongoQueryBuilder) Where(condition string, args ...interface{}) QueryBuilder {
	// for simplicity, just assume "field=?" and similar basic operations
	// example: "age=?", args: 30 → bson.M{"age": 30}
	var field string
	fmt.Sscanf(condition, "%s=?", &field)
	qb.filter[field] = args[0]
	qb.args = append(qb.args, args...)
	return qb
}

func (qb *mongoQueryBuilder) Build() (string, []interface{}) {
	query := bson.M{
		"collection": qb.collection,
		"filter":     qb.filter,
	}
	if len(qb.projection) > 0 {
		query["projection"] = qb.projection
	}
	if len(qb.joins) > 0 {
		query["joins"] = qb.joins
	}
	return fmt.Sprintf("%v", query), qb.args
}

func (db *MongoDB) QueryWithBuilder(ctx context.Context, builder QueryBuilder, result interface{}) error {
	mqb, ok := builder.(*mongoQueryBuilder)
	if !ok {
		return fmt.Errorf("invalid builder type")
	}

	collection := db.Db.Collection(mqb.collection)

	// If joins are used, switch to aggregation
	if len(mqb.joins) > 0 {
		pipeline := make([]bson.M, 0)

		if len(mqb.filter) > 0 {
			pipeline = append(pipeline, bson.M{"$match": mqb.filter})
		}
		pipeline = append(pipeline, mqb.joins...)

		if len(mqb.projection) > 0 {
			pipeline = append(pipeline, bson.M{"$project": mqb.projection})
		}

		cursor, err := collection.Aggregate(ctx, pipeline)
		if err != nil {
			return err
		}
		defer cursor.Close(ctx)

		return cursor.All(ctx, result)
	}

	// Simple find query
	findOpts := options.Find()
	if len(mqb.projection) > 0 {
		findOpts.SetProjection(mqb.projection)
	}

	cursor, err := collection.Find(ctx, mqb.filter, findOpts)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)

	return cursor.All(ctx, result)
}
func (db *MongoDB) Count(ctx context.Context, collection string, filter interface{}) (int64, error) {
	count, err := db.Db.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
