package lib

import (
	"context"
)

// DBInterface is a generic interface for all database operations.
type DBInterface interface {
	Create(ctx context.Context, collection string, data interface{}) error
	Find(ctx context.Context, collection string, query interface{}, result interface{}) error
	Update(ctx context.Context, collection string, query interface{}, updateData interface{}) error
	Delete(ctx context.Context, collection string, query interface{}) error
	// NewQueryBuilder(builderType ...string) QueryBuilder
	QueryWithBuilder(ctx context.Context, builder QueryBuilder, result interface{}) error
}
