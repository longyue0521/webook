package dao

import (
	"github.com/ego-component/egorm"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func InitTables(db *egorm.Component) error {
	// 注册回掉
	err := db.Use(&UserPlugin{})
	if err != nil && !errors.Is(err, gorm.ErrRegistered) {
		return err
	}
	return db.AutoMigrate(
		&User{},
		&UsersIelts{},
	)
}

type UsersIelts User

func (u *UsersIelts) TableName() string {
	return "users_ielts"
}
