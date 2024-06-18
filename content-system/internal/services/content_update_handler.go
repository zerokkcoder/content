package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/dao"
	"github.com/zerokkcoder/content-system/internal/model"
)

type ContentUpdateReq struct {
	ID             int64         `json:"id" binding:"required"`
	Title          string        `json:"title" binding:"required"`
	VideoURL       string        `json:"video_url" binding:"required"`
	Author         string        `json:"author" binding:"required"`
	Description    string        `json:"description"`
	Thumbnail      string        `json:"thumbnail"`
	Category       string        `json:"category"`
	Duration       time.Duration `json:"duration"`
	Resolution     string        `json:"resolution"`
	FileSize       int64         `json:"file_size"`
	Format         string        `json:"format"`
	Quality        int           `json:"quality"`
	ApprovalStatus int           `json:"approval_status"`
}

type ContentUpdateRsp struct {
	Message string `json:"message"`
}

func (ca *CmsApp) ContentUpdate(c *gin.Context) {
	var req ContentUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	contentDao := dao.NewContentDao(ca.db)
	// 判断是否存在
	isExist, err := contentDao.IsExist(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if !isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "内容不存在",
		})
		return
	}
	// 更新
	if err := contentDao.Update(&model.ContentDetail{
		ID:             req.ID,
		Title:          req.Title,
		Description:    req.Description,
		Author:         req.Author,
		VideoURL:       req.VideoURL,
		Thumbnail:      req.Thumbnail,
		Category:       req.Category,
		Duration:       req.Duration,
		Resolution:     req.Resolution,
		FileSize:       req.FileSize,
		Format:         req.Format,
		Quality:        req.Quality,
		ApprovalStatus: req.ApprovalStatus,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": ContentUpdateRsp{Message: "ok"},
	})
}
