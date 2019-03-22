package repository

import (
	"context"

	"github.com/abyssparanoia/merlin/src/model"
	"github.com/jinzhu/gorm"
)

type user struct {
	sql *gorm.DB
}

func (r *user) Get(ctx context.Context, userID int64) (*model.User, error) {

	user := &model.User{}
	user.ID = userID
	r.sql.First(&user)
	return user, nil
}

// NewUser ... ユーザーレポジトリを取得する
func NewUser(sql *gorm.DB) User {
	return &user{
		sql: sql,
	}
}
