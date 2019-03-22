package service

import (
	"context"

	"github.com/abyssparanoia/merlin/src/model"
)

// User ... ユーザーサービスのインターフェイス
type User interface {
	Get(ctx context.Context, userID int64) (*model.User, error)
}
