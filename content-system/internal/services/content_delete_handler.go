package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/api/operate"
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

	rsp, err := ca.operationAppClient.DeleteContent(c, &operate.DeleteContentReq{Id: req.ID})
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
