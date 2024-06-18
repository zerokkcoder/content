package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zerokkcoder/content-system/internal/api"
)

func main() {
	r := gin.Default()
	api.CmsRouters(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	err := r.Run()
	if err != nil {
		fmt.Printf("r run error: %v", err)
		return
	}
}
