package lib

import "os"

type QueryBuilder interface {
	From(table string) QueryBuilder
	Select(fields ...string) QueryBuilder
	Join(join string, condition string) QueryBuilder
	Where(condition string, args ...interface{}) QueryBuilder
	Build() (string, []interface{})
}

func NewQueryBuilder(builderType ...string) QueryBuilder {

	bt := os.Getenv("DB_CONNECTION") // default
	if len(builderType) > 0 {
		bt = builderType[0]
	}

	switch bt {
	case "mysql", "postgres", "sqlite":
		return NewSQLQueryBuilder()
	case "mongodb":
		return NewMongoQueryBuilder()
	default:
		panic("unknown builder type: " + bt)
	}
}
