package tables

import (
	"database/sql"
	"time"
)

type LogicTime struct {
	LogicTimePK int `json:"logic_time_pk" gorm:"primary_key;column:logic_time_pk;comment:Primary key of table logic_time."`

	Value uint64 `json:"value" gorm:"column:value;comment:Logic time value."`

	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at;type:timestamp without time zone;comment:Creation time."`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at;type:timestamp without time zone;comment:Update time."`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:deleted_at;type:timestamp without time zone;comment:Delete time."`
}

func NewLogicTime(lt uint64) *LogicTime {
	now := time.Now().UTC()

	return &LogicTime{
		Value:     lt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (LogicTime) TableName() string {
	return "ton.logic_time"
}
