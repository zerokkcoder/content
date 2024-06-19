package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/zerokkcoder/content-system/internal/dao"
	"github.com/zerokkcoder/content-system/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type LoginReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginRsp struct {
	SessionID string `json:"session_id"`
	Username  string `json:"username"`
	Nickname  string `json:"nickname"`
}

func (ca *CmsApp) Login(c *gin.Context) {
	span := opentracing.SpanFromContext(c.Request.Context())
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	span.SetTag("req", req)
	accountDao := dao.NewAccountDao(ca.db)
	account, err := accountDao.FirstByUsername(c.Request.Context(), req.Username)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "账号不存在",
		})
		return
	}
	// 密码校验
	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "密码错误",
		})
		return
	}
	// 登录成功
	// 生成session
	sessionID, err := ca.genetateSessionID(c.Request.Context(), account.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "系统错误，请稍后重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": LoginRsp{
			SessionID: sessionID,
			Username:  account.Username,
			Nickname:  account.Nickname,
		},
	})
}

func (ca *CmsApp) genetateSessionID(ctx context.Context, username string) (string, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "genetateSessionID")
	defer span.Finish()

	// session id 生成
	sessionID := uuid.New().String()
	// key: session_id:{username} val: session_id 20s
	sessionKey := utils.GetSessionKey(username)
	// session id 持久化
	err := ca.rdb.Set(ctx, sessionKey, sessionID, time.Hour*8).Err()
	if err != nil {
		fmt.Printf("set redis session error = %v\n", err)
		return "", err
	}
	// session id 过期时间 持久化
	authKey := utils.GetAuthKey(sessionID)
	err = ca.rdb.Set(ctx, authKey, time.Now().Unix(), time.Hour*8).Err()
	if err != nil {
		fmt.Printf("set redis auth error = %v\n", err)
		return "", err
	}
	fmt.Println("sessionKey", sessionKey)
	fmt.Println("authKey", authKey)
	return sessionID, nil
}
