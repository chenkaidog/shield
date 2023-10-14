package utils

import (
	"github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
)

func IsEntryDuplicateErr(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == 1062 {
			return true
		}
	}

	if sqliteErr, ok := err.(sqlite3.Error); ok {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return true
		}
	}

	return false
}
