package repository

import (
	"errors"

	"gorm.io/gorm"
)

type TransferRepo interface {
	TransferBalanceAtomic(fromID, toID uint, amount int) error
}

type transferRepo struct {
	db *gorm.DB
}

func NewTransferRepo(db *gorm.DB) TransferRepo {
	return &transferRepo{db: db}
}

func (r transferRepo) TransferBalanceAtomic(fromID, toID uint, amount int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// ลบยอดผู้ส่ง
		result := tx.Exec("UPDATE users SET balance = balance - ? WHERE id = ?", amount, fromID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("sender not found or balance update failed")
		}

		// เพิ่มยอดผู้รับ
		result = tx.Exec("UPDATE users SET balance = balance + ? WHERE id = ?", amount, toID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("receiver not found or balance update failed")
		}

		return nil
	})
}
