package services

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingReq struct {
	Name string `json:"name" binding:"required"`
}

type PingRsp struct {
	Message string `json:"message"`
}

func (ca *CmsApp) Ping(c *gin.Context) {
	var req PingReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": PingRsp{Message: fmt.Sprintf("hello, %s", req.Name)},
	})
}
