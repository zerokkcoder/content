package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/dao"
)

type ContentDeleteReq struct {
	ID int64 `json:"id" binding:"required"`
}

type ContentDeleteRsp struct {
	Message string `json:"message"`
}

func (ca *CmsApp) ContentDelete(c *gin.Context) {
	var req ContentDeleteReq
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
	// 删除
	if err := contentDao.Delete(req.ID) ; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": ContentDeleteRsp{Message: "ok"},
	})
}
