package services

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/dao"
	"github.com/zerokkcoder/content-system/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type RegisterReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

type RegisterRsp struct {
	Message string `json:"message"`
}

func (ca *CmsApp) Register(c *gin.Context) {
	var req RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// 密码加密
	// $2a$10$GiXgbsUv.3TS5Innkg6FgO0AfzDKfsdzpZiZjUZBd2/XaDCN/tI.y
	// $2a$10$lCim7/WvwQN7hCjRdVAVyeDkWiIVd//LvZJ3N69Zy6z39ULxsnGhK
	hashedPassword, err := encryptPassword(req.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{})
	}
	// 账号校验(账号是否存在)
	accountDao := dao.NewAccountDao(ca.db)
	isExist, err := accountDao.IsExist(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if isExist {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "账号已存在",
		})
		return
	}
	// 账号信息持久化
	nowTime := time.Now()
	if err := accountDao.Create(&model.Account{
		Username:  req.Username,
		Password:  hashedPassword,
		Nickname:  req.Nickname,
		CreatedAt: nowTime,
		UpdatedAt: nowTime,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Printf("register req = %+v, hashedPassword = [%s]\n", req, hashedPassword)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": RegisterRsp{Message: "注册成功"},
	})
}

func encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("bcrypt generate from password error = %v\n", err)
		return "", err
	}
	return string(hashedPassword), nil
}
