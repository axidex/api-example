package tables

import (
	"database/sql"
	"time"
)

type Transaction struct {
	TransactionPK int `json:"transaction_pk" gorm:"primary_key;column:transaction_pk;comment:Primary key of table transactions."`

	Source string `json:"source" gorm:"column:source;comment:Source of table transactions."`
	Amount uint64 `json:"amount" gorm:"column:amount;comment:Amount of transactions."`

	CreatedAt time.Time    `json:"created_at" gorm:"column:created_at;type:timestamp without time zone;comment:Creation time."`
	UpdatedAt time.Time    `json:"updated_at" gorm:"column:updated_at;type:timestamp without time zone;comment:Update time."`
	DeletedAt sql.NullTime `json:"deleted_at" gorm:"column:deleted_at;type:timestamp without time zone;comment:Delete time."`
}

func NewTransaction(source string, amount uint64) *Transaction {
	now := time.Now().UTC()

	return &Transaction{
		Source:    source,
		Amount:    amount,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (Transaction) TableName() string {
	return "ton.transactions"
}
