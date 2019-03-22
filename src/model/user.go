package model

// User ... ユーザーモデル
type User struct {
	ID        int64  `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	Sex       string `db:"sex" json:"sex"`
	Enabled   bool   `db:"enabled" json:"enabled"`
	CreatedAt int64  `db:"created_at" json:"created_at"`
	UpdatedAt int64  `db:"updated_at" json:"updated_at"`
}
