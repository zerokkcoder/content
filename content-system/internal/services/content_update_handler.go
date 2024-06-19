package services

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/api/operate"
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
	Quality        int32         `json:"quality"`
	ApprovalStatus int32         `json:"approval_status"`
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

	// 远程调用
	rsp, err := ca.operationAppClient.UpdateContent(c, &operate.UpdateContentReq{
		Content: &operate.Content{
			Id:             req.ID,
			Title:          req.Title,
			Description:    req.Description,
			Author:         req.Author,
			VideoUrl:       req.VideoURL,
			Thumbnail:      req.Thumbnail,
			Category:       req.Category,
			Duration:       req.Duration.Microseconds(),
			Resolution:     req.Resolution,
			FileSize:       req.FileSize,
			Format:         req.Format,
			Quality:        req.Quality,
			ApprovalStatus: req.ApprovalStatus,
		},
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": rsp,
	})
}
