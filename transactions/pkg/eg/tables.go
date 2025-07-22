package eg

import (
	"database/sql"
	"time"
)

type User struct {
	UserID      int     `json:"id_user" gorm:"primaryKey;column:id_user;comment:Primary key of table users."`
	ChatID      int64   `json:"chat_id" gorm:"column:chat_id;unique;not null;comment:Unique chat ID."`
	Username    *string `json:"username" gorm:"column:username;type:varchar(50);comment:Telegram username."`
	StarsCount  int     `json:"stars_count" gorm:"column:stars_count;default:0;comment:Count of stars."`
	TotalEarned int     `json:"total_earned" gorm:"column:total_earned;default:0;comment:Total earned amount."`
}

func (User) TableName() string {
	return "users"
}

type HistoryDeposit struct {
	DepositID    int           `json:"id_deposit" gorm:"primaryKey;column:id_deposit;comment:Primary key of history_deposit."`
	UserID       int           `json:"user_id" gorm:"column:user_id;not null;comment:Foreign key to users."`
	Price        int           `json:"price" gorm:"column:price;not null;comment:Deposit price."`
	Date         time.Time     `json:"date" gorm:"column:date;type:timestamp without time zone;default:CURRENT_TIMESTAMP;comment:Timestamp of deposit."`
	Source       *string       `json:"source" gorm:"column:source;type:varchar(20);comment:Source of deposit."`
	GiftNumberID sql.NullInt64 `json:"id_gift_number" gorm:"column:id_gift_number;comment:Optional gift number."`

	User User `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
}

func (HistoryDeposit) TableName() string {
	return "history_deposit"
}
