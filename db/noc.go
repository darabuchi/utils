//go:build !cgo
// +build !cgo

package db

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func newSqlite(c Config) gorm.Dialector {
	return sqlite.Open(c.Dsn)
}
