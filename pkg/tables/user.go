package tables

import (
	"database/sql"
	"time"
)

type User struct {
	UserPK int `json:"user_pk" gorm:"primary_key;column:user_pk;comment:Primary key of table user."`

	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at;type:timestamp without time zone;comment:Creation time."`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at;type:timestamp without time zone;comment:Update time."`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:deleted_at;type:timestamp without time zone;comment:Delete time."`
}

func (User) TableName() string {
	return "example.user"
}
