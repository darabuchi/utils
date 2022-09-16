package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func getDb() *gorm.DB {
	return _db.Session(&gorm.Session{
		NewDB:             true,
		AllowGlobalUpdate: true,
	})
}

func Select(query interface{}, args ...interface{}) *gorm.DB {
	return getDb().Select(query, args...)
}

func Where(query interface{}, args ...interface{}) *gorm.DB {
	return getDb().Where(query, args...)
}

func Create(value interface{}) *gorm.DB {
	return getDb().Create(value)
}

func Model(model interface{}) *gorm.DB {
	return getDb().Model(model)
}

func Clauses(conds ...clause.Expression) *gorm.DB {
	return getDb().Clauses(conds...)
}

func Unscoped() *gorm.DB {
	return getDb().Unscoped()
}

func Save(value interface{}) *gorm.DB {
	return getDb().Save(value)
}

func Order(value interface{}) *gorm.DB {
	return getDb().Order(value)
}

func Begin() *gorm.DB {
	return getDb().Begin()
}

func Transaction(logic func(tx *gorm.DB) error) error {
	tx := Begin()
	err := logic(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func CommitOrRollback(tx *gorm.DB, logic func() error) error {
	err := logic()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func Statement() *gorm.Statement {
	return getDb().Statement
}

func FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB) {
	return getDb().FirstOrCreate(dest, conds...)
}

func CreateInBatches(value interface{}, batchSize int) *gorm.DB {
	return getDb().CreateInBatches(value, batchSize)
}
