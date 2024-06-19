package data

import (
	"content_manage/internal/biz"
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

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
	ID             int64         `gorm:"column:id;primaryKey"`   // 内容ID
	Title          string        `gorm:"column:title"`           // 内容标题
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

func (*ContentDetail) TableName() string {
	return "cms_content.t_content_details"
}

func (c *contentRepo) Create(ctx context.Context, content *biz.Content) (int64, error) {
	c.log.Infof("contentRepo Create content = %+v", content)
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
	db := c.data.db
	if err := db.Create(&detail).Error; err != nil {
		c.log.Errorf("content create error = %v\n", err)
		return 0, err
	}
	return detail.ID, nil
}

func (c *contentRepo) Update(ctx context.Context, id int64, content *biz.Content) error {
	c.log.Infof("contentRepo Update content = %+v", content)
	db := c.data.db
	detail := ContentDetail{
		ID:             content.ID,
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
	if err := db.Where("id = ?", id).Updates(&detail).Error; err != nil {
		c.log.Errorf("content update error = %v\n", err)
		return err
	}
	return nil
}

func (c *contentRepo) IsExist(ctx context.Context, contentID int64) (bool, error) {
	db := c.data.db
	var detail ContentDetail
	err := db.Where("id = ?", contentID).First(&detail).Error
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
	if err := db.Where("id = ?", id).Delete(&ContentDetail{}).Error; err != nil {
		c.log.WithContext(ctx).Errorf("ContentDao Delete error = %v\n", err)
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
