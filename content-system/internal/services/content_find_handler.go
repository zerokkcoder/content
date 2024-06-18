package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/api/operate"
)

type ContentFindReq struct {
	ID       int64  `json:"id"`
	Author   string `json:"author"`
	Title    string `json:"title"`
	Page     int32  `json:"page"`
	PageSize int32  `json:"page_size"`
}

func (ca *CmsApp) ContentFind(c *gin.Context) {
	var req ContentFindReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// 访问操作服务
	rsp, err := ca.operationAppClient.FindContent(c, &operate.FindContentReq{
		Id:       req.ID,
		Author:   req.Author,
		Title:    req.Title,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
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
