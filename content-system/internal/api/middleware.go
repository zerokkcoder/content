package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/zerokkcoder/content-system/internal/utils"
)

const SessionKey = "session_id"

type SessionAuth struct {
	rdb *redis.Client
}

func NewSessionAuth() *SessionAuth {
	s := &SessionAuth{}
	s.connRdb()
	return s
}

func (s *SessionAuth) Auth(c *gin.Context) {
	sessionID := c.GetHeader(SessionKey)
	// imp auth
	if sessionID == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "session id is null")
	}
	authKey := utils.GetAuthKey(sessionID)
	loginTime, err := s.rdb.Get(c, authKey).Result()
	if err != nil && err != redis.Nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "session auth error")
	}
	if loginTime == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "session auth failed")
	}
	c.Next()
}

func (s *SessionAuth) connRdb() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	s.rdb = rdb
}
