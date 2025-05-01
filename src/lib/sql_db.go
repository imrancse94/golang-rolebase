package lib

import (
	"context"

	"gorm.io/gorm"
)

// SQLDB implements DBInterface for SQL databases
type SQLDB struct {
	Db *gorm.DB
}

func (db *SQLDB) Create(ctx context.Context, collection string, data interface{}) error {
	return db.Db.Table(collection).Create(data).Error
}

func (db *SQLDB) Find(ctx context.Context, collection string, query interface{}, result interface{}) error {
	return db.Db.Table(collection).Where(query).First(result).Error
}

func (db *SQLDB) Update(ctx context.Context, collection string, query interface{}, updateData interface{}) error {
	return db.Db.Table(collection).Where(query).Updates(updateData).Error
}

func (db *SQLDB) Delete(ctx context.Context, collection string, query interface{}) error {
	return db.Db.Table(collection).Where(query).Delete(nil).Error
}
