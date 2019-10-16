package database

import (
	"database/sql"

	otto_runtime "github.com/imiskolee/otto-runtime"
	"github.com/jinzhu/gorm"
)

var resolver func(dbName string) *gorm.DB

func RegisterDatabaseResolver(r func(dbName string) *gorm.DB) {
	resolver = r
}

func init() {
	otto_runtime.Register("NewDatabaseConnection",NewDatabaseConnection)
}


type Database struct {
	db         *gorm.DB
	lastResult sql.Result
}

func NewDatabaseConnection(dbName string) *Database {
	if resolver == nil {
		panic("please setup connection resolver first. ")
	}
	return &Database{
		db: resolver(dbName),
	}
}

func (c *Database) Exec(sql string, vals []interface{}) {
	ret, err := c.db.CommonDB().Exec(sql, vals...)
	if err != nil {
		panic(err)
	}
	c.lastResult = ret
}

func (c *Database) LastInsertID() int64 {
	i, err := c.lastResult.LastInsertId()
	if err != nil {
		panic(err)
	}
	return i
}

func (c *Database) RowsAffected() int64 {
	i, err := c.lastResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	return i
}

/// using map[string]interface for wrapper database columns.
func (c *Database) Query(sql string, vals []interface{}) interface{} {
	rows, err := c.db.CommonDB().Query(sql, vals...)
	if err != nil {
		panic(err)
	}
	cols, _ := rows.Columns()
	var ret []map[string]interface{}
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			panic(err)
		}
		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		ret = append(ret, m)
	}
	return ret
}

func (c *Database) Begin() {
	c.db = c.db.Begin()
}

func (c *Database) Commit() error {
	return c.db.Commit().Error
}

func (c *Database) Rollback() error {
	return c.db.Rollback().Error
}




