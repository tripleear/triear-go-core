package service

import (
	"fmt"

	"github.com/tripleear/triear-go-admin-core/logger"
	"github.com/tripleear/triear-go-admin-core/storage"
	"gorm.io/gorm"
)

type Service struct {
	Orm   *gorm.DB
	Msg   string
	MsgID string
	Log   *logger.Helper
	Error error
	Cache storage.AdapterCache
}

func (db *Service) AddError(err error) error {
	if db.Error == nil {
		db.Error = err
	} else if err != nil {
		db.Error = fmt.Errorf("%v; %w", db.Error, err)
	}
	return db.Error
}

func (db *Service) RunInTransaction(fn func(tx *gorm.DB) error) error {
	// 开始事务
	return db.Orm.Transaction(func(tx *gorm.DB) error {
		// 捕获 panic，确保即使发生 panic 也能回滚事务
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback() // panic 时回滚
				panic(r)      // 继续抛出 panic
			}
		}()

		// 执行传入的操作
		err := fn(tx)
		if err != nil {
			return err
		}
		return nil
	})
}
