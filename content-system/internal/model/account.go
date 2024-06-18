package model

import "time"

type Account struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Nickname  string    `gorm:"column:nickname"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (*Account) TableName() string {
	return "cms_account.account"
}
