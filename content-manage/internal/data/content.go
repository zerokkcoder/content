package data

import (
	"content_manage/internal/biz"
	"context"
	"fmt"
	"hash/fnv"
	"math/big"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

const contentNumTables = 4

type contentRepo struct {
	data *Data
	log  *log.Helper
}

// NewContentRepo .
func NewContentRepo(data *Data, logger log.Logger) biz.ContentRepo {
	return &contentRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type ContentDetail struct {
	ID             int64         `gorm:"column:id;primaryKey"` // 内容ID
	Title          string        `gorm:"column:title"`         // 内容标题
	ContentID      string        `gorm:"column:content_id"`
	Description    string        `gorm:"column:description"`     // 内容描述
	Author         string        `gorm:"column:author"`          // 作者
	VideoURL       string        `gorm:"column:video_url"`       // 视频链接
	Thumbnail      string        `gorm:"column:thumbnail"`       // 封面图
	Category       string        `gorm:"column:category"`        // 内容分类
	Duration       time.Duration `gorm:"column:duration"`        // 内容时长
	Resolution     string        `gorm:"column:resolution"`      // 分辨率 如 720p 1080p
	FileSize       int64         `gorm:"column:file_size"`       // 文件大小
	Format         string        `gorm:"column:format"`          // 文件格式, 如 mp4 avi
	Quality        int32         `gorm:"column:quality"`         // 视频质量 1-高清 2-标清 3-流畅
	ApprovalStatus int32         `gorm:"column:approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	CreatedAt      time.Time     `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      // 更新时间
}

// func (*ContentDetail) TableName() string {
// 	return "cms_content.t_content_details"
// }

type IdxContentDetail struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	ContentID string    `gorm:"column:content_id"`
	Title     string    `gorm:"column:title"`
	Author    string    `gorm:"column:author"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (c *IdxContentDetail) TableName() string {
	return "cms_content.t_idx_content_details"
}

func getContentDetailTable(contentID string) string {
	tableIndex := getContentTableIndex(contentID)
	table := fmt.Sprintf("cms_content.t_content_details_%d", tableIndex)
	log.Infof("content_id = %s, table = %s", contentID, table)
	return table
}

func getContentTableIndex(uuid string) int {
	hash := fnv.New64()
	hash.Write([]byte(uuid))
	hashValue := hash.Sum64()
	fmt.Println("hashValue = ", hashValue)

	bigNum := big.NewInt(int64(hashValue))
	mod := big.NewInt(contentNumTables)
	tableIndex := bigNum.Mod(bigNum, mod).Int64()
	return int(tableIndex)
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) (int64, error) {
	c.log.Infof("contentRepo Create content = %+v", content)
	db := c.data.db

	idx := IdxContentDetail{
		ContentID: content.ContentID,
		Title:     content.Title,
		Author:    content.Author,
	}
	if err := db.Create(&idx).Error; err != nil {
		c.log.Errorf("content create error = %v\n", err)
		return 0, err
	}

	detail := ContentDetail{
		Title:          content.Title,
		ContentID:      content.ContentID,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}

	if err := db.Table(getContentDetailTable(content.ContentID)).Create(&detail).Error; err != nil {
		c.log.Errorf("content create error = %v\n", err)
		return 0, err
	}
	return idx.ID, nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	c.log.Infof("contentRepo Update content = %+v", content)
	db := c.data.db
	var idx IdxContentDetail
	if err := db.Where("id = ?", id).First(&idx).Error; err != nil {
		return err
	}

	detail := ContentDetail{
		Title:          content.Title,
		Description:    content.Description,
		Author:         content.Author,
		VideoURL:       content.VideoURL,
		Thumbnail:      content.Thumbnail,
		Category:       content.Category,
		Duration:       content.Duration,
		Resolution:     content.Resolution,
		FileSize:       content.FileSize,
		Format:         content.Format,
		Quality:        content.Quality,
		ApprovalStatus: content.ApprovalStatus,
	}
	if err := db.Table(getContentDetailTable(idx.ContentID)).Where("content_id = ?", idx.ContentID).Updates(&detail).Error; err != nil {
		c.log.Errorf("content update error = %v\n", err)
		return err
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, id int64) (bool, error) {
	db := c.data.db
	var detail IdxContentDetail
	err := db.Where("id = ?", id).First(&detail).Error
	if err == gorm.ErrRecordNotFound {
		return false, nil
	}
	if err != nil {
		c.log.WithContext(ctx).Errorf("contentRepo IsExist error = %v\n", err)
		return false, err
	}
	return true, nil
}

func (c *contentRepo) Delete(ctx context.Context, id int64) error {
	db := c.data.db
	// 查询索引表信息
	var idx IdxContentDetail
	if err := db.Where("id = ?", id).First(&idx).Error; err != nil {
		return err
	}
	// 删除索引信息
	if err := db.Where("id = ?", id).Delete(&IdxContentDetail{}).Error; err != nil {
		c.log.WithContext(ctx).Errorf("ContentDao IdxContentDetail Delete error = %v\n", err)
		return err
	}
	// 删除详情信息
	if err := db.Table(getContentDetailTable(idx.ContentID)).
		Where("content_id = ?", idx.ContentID).Delete(&ContentDetail{}).Error; err != nil {
		c.log.WithContext(ctx).Errorf("ContentDao ContentDetail Delete error = %v\n", err)
		return err
	}
	return nil
}

func (c *contentRepo) Find(ctx context.Context, params *biz.FindParams) ([]*biz.Content, int64, error) {
	query := c.data.db.Model(&ContentDetail{})
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
		return nil, 0, err
	}

	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = int(params.Page)
	}
	if params.PageSize > 0 {
		pageSize = int(params.PageSize)
	}
	offset := (page - 1) * pageSize
	var results []*ContentDetail
	if err := query.Offset(offset).
		Limit(pageSize).
		Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("contentRepo Find error = %v\n", err)
		return nil, 0, err
	}

	contents := make([]*biz.Content, 0, len(results))
	for _, v := range results {
		contents = append(contents, &biz.Content{
			ID:             v.ID,
			Title:          v.Title,
			VideoURL:       v.VideoURL,
			Author:         v.Author,
			Description:    v.Description,
			Thumbnail:      v.Thumbnail,
			Category:       v.Category,
			Duration:       v.Duration,
			Resolution:     v.Resolution,
			FileSize:       v.FileSize,
			Format:         v.Format,
			Quality:        v.Quality,
			ApprovalStatus: v.ApprovalStatus,
			UpdatedAt:      v.UpdatedAt,
			CreatedAt:      v.CreatedAt,
		})
	}

	return contents, total, nil
}

func (c *contentRepo) FindIndex(ctx context.Context, params *biz.FindParams) ([]*biz.ContentIndex, int64, error) {
	// 构建查询条件
	query := c.data.db.Model(&IdxContentDetail{})
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
		return nil, 0, err
	}

	var page, pageSize = 1, 10
	if params.Page > 0 {
		page = int(params.Page)
	}
	if params.PageSize > 0 {
		pageSize = int(params.PageSize)
	}
	offset := (page - 1) * pageSize
	var results []*IdxContentDetail
	if err := query.Offset(offset).
		Limit(pageSize).
		Find(&results).Error; err != nil {
		c.log.WithContext(ctx).Errorf("contentRepo FindIndex error = %v\n", err)
		return nil, 0, err
	}

	var contents []*biz.ContentIndex
	for _, v := range results {
		contents = append(contents, &biz.ContentIndex{
			ID:        v.ID,
			ContentID: v.ContentID,
		})
	}

	c.log.Infof("contentRepo FindIndex content = %+v", contents)

	return contents, total, nil
}

func (c *contentRepo) First(ctx context.Context, idx *biz.ContentIndex) (*biz.Content, error) {
	db := c.data.db
	var detail ContentDetail
	c.log.Infof("contentRepo First ContentID = %s", idx.ContentID)
	if err := db.Table(getContentDetailTable(idx.ContentID)).
		Where("content_id = ?", idx.ContentID).First(&detail).Error; err != nil {
		c.log.WithContext(ctx).Errorf("contentRepo First error = %v\n", err)
		return nil, err
	}

	content := &biz.Content{
		ID:             idx.ID,
		ContentID:      idx.ContentID,
		Title:          detail.Title,
		VideoURL:       detail.VideoURL,
		Author:         detail.Author,
		Description:    detail.Description,
		Thumbnail:      detail.Thumbnail,
		Category:       detail.Category,
		Duration:       detail.Duration,
		Resolution:     detail.Resolution,
		FileSize:       detail.FileSize,
		Format:         detail.Format,
		Quality:        detail.Quality,
		ApprovalStatus: detail.ApprovalStatus,
		UpdatedAt:      detail.UpdatedAt,
		CreatedAt:      detail.CreatedAt,
	}
	return content, nil
}
