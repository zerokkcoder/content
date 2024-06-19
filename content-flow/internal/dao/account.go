package dao

import (
	"fmt"

	"github.com/zerokkcoder/content-flow/internal/model"
	"gorm.io/gorm"
)

type AccountDao struct {
	db *gorm.DB
}

func NewAccountDao(db *gorm.DB) *AccountDao {
	return &AccountDao{db: db}
}

func (a *AccountDao) IsExist(username string) (bool, error) {
	var account model.Account
	err := a.db.Where("username = ?", username).First(&account).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("AccountDao IsExist error = %v\n", err)
		return false, err
	}
	return true, nil
}

func (a *AccountDao) Create(account *model.Account) error {
	if err := a.db.Create(account).Error; err != nil {
		fmt.Printf("AccountDao Create error = %v\n", err)
		return err
	}
	return nil
}

func (a *AccountDao) FirstByUsername(username string) (*model.Account, error) {
	var account model.Account
	if err := a.db.Where("username = ?", username).First(&account).Error; err != nil {
		fmt.Printf("AccountDao FirstByUsername error = %v\n", err)
		return nil, err
	}
	return &account, nil
}
