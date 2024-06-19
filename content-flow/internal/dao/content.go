package dao

import (
	"fmt"
	"log"

	"github.com/zerokkcoder/content-flow/internal/model"
	"gorm.io/gorm"
)

type ContentDao struct {
	db *gorm.DB
}

func NewContentDao(db *gorm.DB) *ContentDao {
	return &ContentDao{db: db}
}

func (c *ContentDao) First(contentID int64) (*model.ContentDetail, error) {
	var detail model.ContentDetail
	if err := c.db.Where("id = ?", contentID).First(&detail).Error; err != nil {
		fmt.Printf("ContentDao First error = %v\n", err)
		return nil, err
	}
	return &detail, nil
}

func (c *ContentDao) Create(detail *model.ContentDetail) (int64, error) {
	if err := c.db.Create(detail).Error; err != nil {
		fmt.Printf("ContentDao Create error = %v\n", err)
		return 0, err
	}
	return detail.ID, nil
}

func (c *ContentDao) Update(detail *model.ContentDetail) error {
	if err := c.db.
		Where("id = ?", detail.ID).
		Updates(detail).Error; err != nil {
		fmt.Printf("ContentDao Update error = %v\n", err)
		return err
	}
	return nil
}

func (c *ContentDao) IsExist(contentID int64) (bool, error) {
	var detail model.ContentDetail
	err := c.db.Where("id = ?", contentID).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		fmt.Printf("ContentDao IsExist error = %v\n", err)
		return false, err
	}
	return true, nil
}

func (c *ContentDao) Delete(contentID int64) error {
	if err := c.db.Where("id = ?", contentID).Delete(&model.ContentDetail{}).Error; err != nil {
		fmt.Printf("ContentDao Delete error = %v\n", err)
		return err
	}
	return nil
}

type FindParams struct {
	ID       int64
	Author   string
	Title    string
	Page     int
	PageSize int
}

func (c *ContentDao) Find(params *FindParams) ([]*model.ContentDetail, int64, error) {
	query := c.db.Model(&model.ContentDetail{})
	if params.ID != 0 {
		query = query.Where("id = ?", params.ID)
	}
	if params.Author != "" {
		query = query.Where("author = ?", params.Author)
	}
	if params.Title != "" {
		query = query.Where("title LIKE ?", "%"+params.Title+"%")
	}
	var total int64
	if err := query.Count(&total).Error; err != nil {
		fmt.Printf("ContentDao Find error = %v\n", err)
		return nil, 0, err
	}

	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = params.Page
	}
	if params.PageSize > 0 {
		pageSize = params.PageSize
	}
	offset := (page - 1) * pageSize
	var data []*model.ContentDetail
	if err := query.Offset(offset).
		Limit(pageSize).
		Find(&data).Error; err != nil {
		fmt.Printf("ContentDao Find error = %v\n", err)
		return nil, 0, err
	}

	return data, total, nil
}

func (c *ContentDao) UpdateByID(id int64, column string, value interface{}) error {
	if err := c.db.Model(&model.ContentDetail{}).
		Where("id = ?", id).
		Update(column, value).Error; err != nil {
		log.Printf("ContentDao UpdateByID error = %v\n", err)
		return err
	}
	return nil
}
