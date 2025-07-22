package eg

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

var sourceName = "TON"

type Storage interface {
	DepositToUser(ctx context.Context, chatId int64, price int) error
}

type StorageGorm struct {
	db *gorm.DB
}

func NewEGStorage(db *gorm.DB) *StorageGorm {
	return &StorageGorm{db: db}
}

func (s *StorageGorm) DepositToUser(ctx context.Context, chatId int64, price int) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var user User
		if err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
			Where("chat_id = ?", chatId).
			First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				user = User{ChatID: chatId}
				if err := tx.Create(&user).Error; err != nil {
					return err
				}
			} else {
				return err
			}
		}

		user.StarsCount += price
		if err := tx.Model(&user).Update("stars_count", user.StarsCount).Error; err != nil {
			return err
		}

		history := HistoryDeposit{
			UserID: user.UserID,
			Price:  price,
			Date:   time.Now().UTC(),
			Source: &sourceName,
		}

		if err := tx.Create(&history).Error; err != nil {
			return err
		}

		return nil
	})
}
