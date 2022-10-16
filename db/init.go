package db

import (
	"errors"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"

	"github.com/darabuchi/log"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var _db *gorm.DB

func Connect(c Config, tables ...interface{}) error {
	if _db != nil {
		return nil
	}

	log.Infof("connecting to database (%s)%s ", c.Database, c.Dsn)

	var d gorm.Dialector
	switch c.Database {
	case MySql:
		d = mysql.New(mysql.Config{
			DSN:                    c.Dsn,
			DefaultStringSize:      500,
			DontSupportRenameIndex: true,
		})
	case Sqlite:
		d = newSqlite(c)
	case Postgres:
		d = postgres.New(postgres.Config{
			DriverName:           c.DriverName,
			DSN:                  c.Dsn,
			PreferSimpleProtocol: true,
			WithoutReturning:     false,
			Conn:                 nil,
		})
	case Sqlserver:
		d = sqlserver.Open(c.Dsn)
	// case Clickhouse:
	// 	d = clickhouse.Open(c.Dsn)
	default:
		return errors.New("unknown database")
	}

	var err error
	_db, err = gorm.Open(d, &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		FullSaveAssociations:                     false,
		Logger:                                   NewLogger(),
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: true,
		DisableNestedTransaction:                 true,
		AllowGlobalUpdate:                        true,
		QueryFields:                              false,
		CreateBatchSize:                          100,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	if c.Debug {
		_db = _db.Debug()
	}

	conn, err := _db.DB()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	err = conn.Ping()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	err = AutoMigrate(tables...)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

func AutoMigrate(dst ...interface{}) error {
	return getDb().AutoMigrate(dst...)
}
