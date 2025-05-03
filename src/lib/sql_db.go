package lib

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// SQLDB implements DBInterface for SQL databases
type SQLDB struct {
	Db *gorm.DB
}

type SQLQueryBuilder struct {
	table   string
	selects []string
	joins   []string
	where   []string
	args    []interface{}
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

// FindAll retrieves all records matching the query in SQL
func (db *SQLDB) FindAll(ctx context.Context, collection string, query interface{}, result interface{}) error {
	return db.Db.Table(collection).Where(query).Find(result).Error
}

func (db *SQLDB) QueryWithBuilder(ctx context.Context, builder QueryBuilder, result interface{}) error {
	query, args := builder.Build()
	return db.Db.WithContext(ctx).Raw(query, args...).Scan(result).Error
}

// QueryBuilder interface for building SQL queries
func NewSQLQueryBuilder() *SQLQueryBuilder {
	return &SQLQueryBuilder{}
}

func (b *SQLQueryBuilder) From(table string) QueryBuilder {
	b.table = table
	return b
}

func (b *SQLQueryBuilder) Select(fields ...string) QueryBuilder {
	b.selects = append(b.selects, fields...)
	return b
}

func (b *SQLQueryBuilder) Join(join string, condition string) QueryBuilder {
	b.joins = append(b.joins, fmt.Sprintf("JOIN %s ON %s", join, condition))
	return b
}

func (b *SQLQueryBuilder) Where(condition string, args ...interface{}) QueryBuilder {
	b.where = append(b.where, condition)
	b.args = append(b.args, args...)
	return b
}

func (b *SQLQueryBuilder) Build() (string, []interface{}) {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(b.selects, ", "), b.table)

	if len(b.joins) > 0 {
		query += " " + strings.Join(b.joins, " ")
	}

	if len(b.where) > 0 {
		query += " WHERE " + strings.Join(b.where, " AND ")
	}

	return query, b.args
}
