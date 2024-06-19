package model

import "time"

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
	Quality        int           `gorm:"column:quality"`         // 视频质量 1-高清 2-标清 3-流畅
	ApprovalStatus int           `gorm:"column:approval_status"` // 审核状态 1-审核中 2-审核通过 3-审核不通过
	CreatedAt      time.Time     `gorm:"column:created_at"`      // 创建时间
	UpdatedAt      time.Time     `gorm:"column:updated_at"`      // 更新时间
}

// func (*ContentDetail) TableName() string {
// 	return "cms_content.t_content_details"
// }
